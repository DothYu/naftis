package storer

import (
	"github.com/xiaomi/naftis/src/api/storer/db"
)

/**
 * description: 初始化storer包
 */
func Init() {
	db.Init()
}
