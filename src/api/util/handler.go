package util

import (
	"errors"
	"net/http"

	"github.com/xiaomi/naftis/src/api/log"
	"github.com/xiaomi/naftis/src/api/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// BindFailFn defines function that handle payload-binding failure.
var BindFailFn = func(c *gin.Context, e error) {
	log.Info("[BindFailFn]", "err", e)
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 1,
		"data": e.Error(),
	})
}

// OpFailFn defines function that handle internal error.
var OpFailFn = func(c *gin.Context, e error) {
	log.Info("[OpFailFn]", "err", e)
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 2,
		"data": e.Error(),
	})
}

/**
 * description: RetOK: 一次成功的HTTP JSON响应
 */
var RetOK = map[string]interface{}{"code": 0, "data": struct{}{}}

const (
	// JWTSecret defines default JWT secret
	JWTSecret = "istioIsAwesome"
)

var (
	// ErrJWTUnauthorized is returned when the token isn't authorized.
	ErrJWTUnauthorized = errors.New("unauthorized")
)

// Authentication authenticates incoming request.
var Authentication = func(r *http.Request) (model.User, error) {
	t, e := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(JWTSecret)
		return b, nil
	})
	if e != nil {
		return model.User{}, ErrJWTUnauthorized
	}

	username := t.Claims.(jwt.MapClaims)["username"].(string)
	u, ok := model.MockUsers[username]
	if !ok {
		return model.User{}, ErrJWTUnauthorized
	}
	return u, nil
}
