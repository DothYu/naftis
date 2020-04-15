package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaomi/naftis/src/api/bootstrap"
	"github.com/xiaomi/naftis/src/api/service"
	"github.com/xiaomi/naftis/src/api/util"
)

/**
 * description: 返回所有可获得的service简要信息
 */
func Services(c *gin.Context) {
	t := c.Query("t")
	if t != "tree" {
		uid := c.Param("uid")
		c.JSON(200, gin.H{
			"code": 0,
			"data": service.ServiceInfo.Services(uid).Exclude("kube-system", bootstrap.Args.IstioNamespace, bootstrap.Args.Namespace).Status(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": service.ServiceInfo.Tree(),
	})
	return
}

/**
 * description: 通过service UID来查询其Pods
 */
func ServicePods(c *gin.Context) {
	uid := c.Param("uid")
	svcs := service.ServiceInfo.Services(uid)

	if len(svcs) == 0 || len(svcs[0].Labels) == 0 {
		c.JSON(200, util.RetOK)
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": svcs[0].Pods.Status(),
	})
}

/**
 * description: 通过pod name来获取pod信息
 */
func Pods(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"code": 0,
		"data": service.ServiceInfo.PodsByName(name).Exclude("kube-system", bootstrap.Args.IstioNamespace, bootstrap.Args.Namespace).Status(),
	})
}

/**
 * description: returns data like namespaces of Kubernetes.
 */
func Kubeinfo(c *gin.Context) {
	var ns = service.ServiceInfo.Namespaces("").Exclude("kube-system", bootstrap.Args.Namespace)
	var retNs = make([]string, 0, len(ns))
	for _, n := range ns {
		retNs = append(retNs, n.Name)
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": map[string]interface{}{
			"namespaces": retNs,
		},
	})
}
