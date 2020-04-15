package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaomi/naftis/src/api/handler"
	"github.com/xiaomi/naftis/src/api/middleware"
)

/**
 * description: 初始化router包
 */
func Init(e *gin.Engine) {
	e.Use(
		gin.Recovery(),
		// middleware.XSS(), // TODO fix XSS middleware leading to JSON Binding bugs.{"varMaps":[{&#34;Host&#34;:&#34;details&#34;,&#34;DestinationSubset&#34;:&#34;v3&#34;}],"command":"apply","tmplID":43,"serviceUID":"182dcf63-9a26-11e8-bd9d-525400c9c704"}
	)

	// 公共 APIs
	e.GET("/api/probe/healthy", handler.Healthy)       // 应用健康检测
	e.POST("/api/login/account", handler.LoginAccount) // 验证用户账号

	// 分组 APIs
	api := e.Group("/api")
	api.Use(middleware.Auth())

	api.GET("/login_user", handler.LoginUser) // 返回当前用户
	api.GET("/diagnose", handler.ListStatus)  // Istio诊断（service、pod简要信息）
	api.GET("/metrics", handler.ListMetrics)  // 返回服务网格的一些概述指标（service数目，pod数目）
	api.GET("/d3graph", handler.D3Graph)      // 返回按照提供的root service name 过滤的d3图表

	api.GET("/services", handler.Services)              // 返回所有services下的service tree
	api.GET("/services/:uid", handler.Services)         // 返回指定service下的service tree
	api.GET("/services/:uid/pods", handler.ServicePods) // 通过service UID来查询其Pods的简要信息

	api.GET("/pods", handler.Pods)       // 返回所有pods信息
	api.GET("/pods/:name", handler.Pods) // 通过pod name来获取pod信息

	api.GET("/tasks", handler.ListTasks)     // 返回所有已保存的任务
	api.GET("/tasks/:id", handler.ListTasks) // 返回指定的已保存任务
	api.POST("/tasks", handler.AddTasks)     // 添加一个任务到执行器中

	api.GET("/tasktmpls", handler.ListTaskTmpls)             // 返回所有任务模板
	api.GET("/tasktmpls/:id", handler.ListTaskTmpls)         // 返回指定的任务模板
	api.POST("/tasktmpls", handler.AddTaskTmpls)             // 新建一个任务模板
	api.PUT("/tasktmpls/:id", handler.UpdateTaskTmpls)       // 更新一个任务模板
	api.DELETE("/tasktmpls/:id", handler.DeleteTaskTmpls)    // 删除一个任务模板
	api.GET("/tasktmpls/:id/vars", handler.ListTaskTmplVars) // 返回指定任务模板的变量map
	api.GET("/kube/info", handler.Kubeinfo)                  // 返回k8s中namespace

	// 构造一个hub实例
	hub := handler.NewHub()
	go hub.Run()
	e.GET("/ws", func(c *gin.Context) {
		handler.ServeWS(hub, c.Writer, c.Request)
	})
}
