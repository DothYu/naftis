package db

import (
	"errors"
	"fmt"

	"github.com/xiaomi/naftis/src/api/bootstrap"
	"github.com/xiaomi/naftis/src/api/log"
	"github.com/xiaomi/naftis/src/api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import mysql driver
	"github.com/spf13/viper"
)

var (
	db *gorm.DB
)

/**
 * description: 迁移数据库表
 */
func migrate() {
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.TaskTmpl{})
}

var (
	// 非法参数错误
	ErrInvalidParams = errors.New("invalid params")
	// sql执行错误
	ErrSQLExec = errors.New("sql executed fail")
)

/**
 * description: 初始化db包
 */
func Init() {
	var err error
	// 数据库连接
	db, err = gorm.Open("mysql", viper.GetString("db.default"))
	if err != nil {
		panic(fmt.Errorf("failed to connect database %s", err))
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(bootstrap.Debug())

	log.Info("[db] init success")
}
