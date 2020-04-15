package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaomi/naftis/src/api/bootstrap"
	"github.com/xiaomi/naftis/src/api/service"
)

/**
 * description: 返回服务网格的一些概述指标（service数目，pod数目）
 */
func ListMetrics(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": map[string]interface{}{
			"serviceCount": len(service.ServiceInfo.Services("").Exclude("kube-system", bootstrap.Args.IstioNamespace, bootstrap.Args.Namespace)),
			"podCount":     len(service.ServiceInfo.Pods().Exclude("kube-system", bootstrap.Args.IstioNamespace, bootstrap.Args.Namespace)),
		},
	})
}
