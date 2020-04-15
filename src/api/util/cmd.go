package util

import (
	"os"
	"path/filepath"

	"github.com/xiaomi/naftis/src/api/bootstrap"
)

/**
 * description: 返回home path
 */
func Home() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

/**
 * description: 返回kube config path
 */
func Kubeconfig() string {
	if bootstrap.Args.InCluster {
		return ""
	}
	return filepath.Join(Home(), ".kube", "config")
}
