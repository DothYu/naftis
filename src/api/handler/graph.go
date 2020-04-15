package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaomi/naftis/src/api/service"
)

/**
 * description: 返回按照提供的root service name 过滤的d3图表
 */
func D3Graph(c *gin.Context) {
	service.Prom.ServeHTTP(c.Writer, c.Request)
	return
}
