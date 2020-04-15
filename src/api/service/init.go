package service

import (
	"github.com/xiaomi/naftis/src/api/storer"
)

/**
 * description: 初始化service包
 */
func Init() {
	storer.Init() // 初始化storer包
	InitKube()    // 初始化kube
	InitProm()    // 初始化prometheus服务
}
