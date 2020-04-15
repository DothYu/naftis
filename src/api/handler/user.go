package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/xiaomi/naftis/src/api/model"
	"github.com/xiaomi/naftis/src/api/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/**
 * description: 返回当前登录用户
 */
func LoginUser(c *gin.Context) {
	ret := struct {
		Name        string `json:"name"`
		Avatar      string `json:"avatar"`
		UserID      string `json:"userid"`
		NotifyCount int    `json:"notifyCount"`
	}{
		util.User(c).Name,
		"assets/mi-black.png",
		"001",
		12,
	}
	c.JSON(200, ret)
}

type accountPayload struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Type     string `json:"type"`
}

func (a *accountPayload) validate() error {
	if a.Password == "" || a.Username == "" {
		return errors.New("empty username or password")
	}
	v, _ := model.MockUsers[a.Username]
	if v.Password != a.Password {
		return errors.New("invalid username or password")
	}
	return nil
}

/**
 * description: 验证用户账号
 */
func LoginAccount(c *gin.Context) {
	var p accountPayload
	if e := c.BindJSON(&p); e != nil {
		util.BindFailFn(c, e)
		return
	}
	if e := p.validate(); e != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": e.Error(),
		})
		return
	}
	// 获取 token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"username": p.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	tokenString, err := token.SignedString([]byte(util.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"data": "Could not generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": map[string]string{
			"status":           "ok",
			"type":             p.Type,
			"currentAuthority": "admin",
			"token":            tokenString,
		},
	})
}
