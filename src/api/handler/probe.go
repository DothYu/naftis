package handler

import (
	"github.com/gin-gonic/gin"
)

/**
 * description: 返回后端的服务健康信息
 */
func Healthy(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": "",
	})
}
