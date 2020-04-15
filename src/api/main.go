package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/xiaomi/naftis/src/api/bootstrap"
	"github.com/xiaomi/naftis/src/api/executor"
	"github.com/xiaomi/naftis/src/api/log"
	"github.com/xiaomi/naftis/src/api/router"
	"github.com/xiaomi/naftis/src/api/service"
	"github.com/xiaomi/naftis/src/api/version"
	"github.com/xiaomi/naftis/src/api/worker"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/**
 * description: Cobra: 三个基本概念commands,arguments和flags
 * commands: 行为
 * arguments: 数值
 * flags: 对行为的改变
 * 例子: # git是appname，clone是commands，URL是arguments，brae是flags
		git clone URL --bare
*/
var (
	rootCmd = &cobra.Command{
		Use:               "naftis-api",
		Short:             "naftis API server",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Long:              `Start naftis API server`,
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	}
	startCmd = &cobra.Command{
		Use:     "start",
		Short:   "Start naftis API server",
		Example: "naftis-api start -c config/in-local.toml.toml",
		RunE:    start,
	}
)

func init() {
	startCmd.PersistentFlags().StringVarP(&bootstrap.Args.ConfigFile, "config", "c", "config/in-local.toml",
		"Start server with provided configuration file")
	startCmd.PersistentFlags().StringVarP(&bootstrap.Args.Host, "host", "H", "0.0.0.0",
		"Start server with provided host")
	startCmd.PersistentFlags().IntVarP(&bootstrap.Args.Port, "port", "p", 50000,
		"Start server with provided port")
	startCmd.PersistentFlags().BoolVarP(&bootstrap.Args.InCluster, "inCluster", "i", true,
		"Start server in kube cluster")
	startCmd.PersistentFlags().StringVarP(&bootstrap.Args.IstioNamespace, "istioNamespace", "I", "istio-system",
		"Start server with provided deployed Istio namespace")

	// 将命令添加到父项中
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(version.Command())
}

/**
 * description: startCmd定义的RunE
 */
func start(_ *cobra.Command, _ []string) error {
	parseConfig()

	log.Init()
	// 读取配置文件中的应用启动模式并设置
	mode := viper.GetString("mode")
	gin.SetMode(mode)
	bootstrap.SetDebug(mode)

	// 监听进程结束的信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 设置后端服务
	engine := gin.Default()
	executor.Init()
	router.Init(engine)
	service.Init()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", bootstrap.Args.Host, bootstrap.Args.Port),
		Handler: engine,
	}

	// 开启任务执行器
	go worker.Start()

	// 优雅关闭服务
	go func() {
		<-quit
		worker.Stop()
		fmt.Println("stopping server now")
		if err := server.Close(); err != nil {
			fmt.Println("Server Close:", err)
		}
	}()

	// 开启后端服务
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			fmt.Printf("Server closed under request\n")
		} else {
			fmt.Printf("Server closed unexpect, %s\n", err.Error())
		}
	}

	return nil
}

/**
 * description: 应用配置文件初始化
 */
func parseConfig() {
	// 使用viper来管理配置文件
	viper.SetConfigFile(bootstrap.Args.ConfigFile)
	// 根据以上配置读取加载配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("parse config file fail: %s", err))
	}

	// 初始化 Naftis namespace
	b, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace") // just pass the file name
	if err != nil || string(b) == "" {
		log.Info("[k8s] get Naftis namespace fail or get empty namespace, use `naftis` by default", "err", err, "namespace", string(b))
		bootstrap.Args.Namespace = "naftis"
	} else {
		bootstrap.Args.Namespace = string(b)
	}
	log.Info("[Args]", "Host", bootstrap.Args.Host)
	log.Info("[Args]", "Port", bootstrap.Args.Port)
	log.Info("[Args]", "InCluster", bootstrap.Args.InCluster)
	log.Info("[Args]", "ConfigFile", bootstrap.Args.ConfigFile)
	log.Info("[Args]", "Namespace", bootstrap.Args.Namespace)
	log.Info("[Args]", "IstioNamespace", bootstrap.Args.IstioNamespace)
	println()
}

func main() {
	// 初始化Cobra
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
