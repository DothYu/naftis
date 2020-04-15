package util

import (
	"github.com/xiaomi/naftis/src/api/model"

	"github.com/gin-gonic/gin"
)

/**
 * description: 返回现在的登录用户
 */
func User(c *gin.Context) model.User {
	user, _ := c.Get("user")
	return user.(model.User)
}
