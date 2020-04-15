package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaomi/naftis/src/api/log"
	"github.com/xiaomi/naftis/src/api/service"
)

/**
 * description: 返回istio中所有service和pod的简要信息
 */
func ListStatus(c *gin.Context) {
	log.Info("[API] /api/diagnose start", "ts", time.Now())

	log.Info("[API] /api/diagnose Services start", "ts", time.Now())
	svcs := service.IstioInfo.Services("").Status()
	log.Info("[API] /api/diagnose Services end", "ts", time.Now())

	log.Info("[API] /api/diagnose Pods start", "ts", time.Now())
	pods := service.IstioInfo.Pods().Status()
	log.Info("[API] /api/diagnose Pods end", "ts", time.Now())

	c.JSON(200, gin.H{
		"code": 0,
		"data": map[string]interface{}{
			"components": svcs,
			"pods":       pods,
		},
	})
	log.Info("[API] /api/diagnose end", "ts", time.Now())
}
