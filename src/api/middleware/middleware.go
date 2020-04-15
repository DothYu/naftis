package middleware

import (
	"net/http"

	"github.com/dvwright/xss-mw"
	"github.com/gin-gonic/gin"
	"github.com/xiaomi/naftis/src/api/util"
)

var (
	xssMw xss.XssMw
)

/**
 * description: 定义JWT认证中间件
 */
var Auth = func() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, e := util.Authentication(c.Request)
		if e != nil {
			c.AbortWithError(http.StatusUnauthorized, e)
			return
		}

		c.Set("user", u)
	}
}

// XSS defines Xss middleware.
var XSS = xssMw.RemoveXss
