package bootstrap

var debug bool

// Args提供了naftis-api服务启动的所有参数
var Args args

type args struct {
	Host           string
	Port           int
	InCluster      bool
	ConfigFile     string
	Namespace      string
	IstioNamespace string
}

/**
 * description: 设置应用启动模式
 */
func SetDebug(mode string) {
	if mode == "debug" {
		debug = true
	}
}

/**
 * description: 返回应用启动模式
 */
func Debug() bool {
	return debug
}
