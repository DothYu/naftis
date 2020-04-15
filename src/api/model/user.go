package model

import (
	"encoding/gob"
)

func init() {
	gob.Register(&User{})
}

/**
 * description: 用户结构体
 */
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	DepID    int    `json:"depId"`
	Email    string `json:"email"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
}

/**
 * description: 模拟用户数据
 */
var MockUsers = map[string]User{
	"admin": {
		ID:       1,
		Name:     "admin",
		Password: "admin",
	},
	"user": {
		ID:       2,
		Name:     "user",
		Password: "user",
	},
	"test-01": {
		ID:       3,
		Name:     "test-01",
		Password: "test-01",
	},
	"test-02": {
		ID:       4,
		Name:     "test-02",
		Password: "test-02",
	},
}
