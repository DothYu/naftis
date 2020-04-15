package handler

import (
	"github.com/xiaomi/naftis/src/api/model"
	"github.com/xiaomi/naftis/src/api/service"
	"github.com/xiaomi/naftis/src/api/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type taskTmplPayload struct {
	ID      uint                `json:"id"`
	Name    string              `json:"name"`
	Command int                 `json:"command"`
	Content string              `json:"content"`
	Brief   string              `json:"brief"`
	Vars    []model.TaskTmplVar `json:"vars"`
	Icon    string              `json:"icon"`
	Default string              `string:"default"`
}

/**
 * description: 返回指定的任务模板
 */
func ListTaskTmpls(c *gin.Context) {
	ids := make([]uint, 0, 1)
	if idStr := c.Param("id"); idStr != "" {
		ids = append(ids, cast.ToUint(idStr))
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": service.TaskTmpl.Get("", "", "", ids, 0, 0, 0, 0),
	})
}

/**
 * description: 返回指定task模板的variable map
 */
func ListTaskTmplVars(c *gin.Context) {
	var taskTmplID = cast.ToUint(c.Param("id"))
	c.JSON(200, gin.H{
		"code": 0,
		"data": service.TaskTmplVar.Get("", "", "", "", 0, taskTmplID, []uint{}),
	})
}

/**
 * description: 新建一个任务模板
 */
func AddTaskTmpls(c *gin.Context) {
	var p taskTmplPayload
	if e := c.BindJSON(&p); e != nil {
		util.BindFailFn(c, e)
		return
	}
	t, e := service.TaskTmpl.Add(p.Name, p.Content, p.Brief, util.User(c).Name, p.Icon, p.Vars)
	if e != nil {
		util.OpFailFn(c, e)
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": []interface{}{t},
	})
}

/**
 * description: 更新一个任务模板
 */
func UpdateTaskTmpls(c *gin.Context) {
	var id = cast.ToUint(c.Param("id"))
	var p taskTmplPayload
	if e := c.BindJSON(&p); e != nil {
		util.BindFailFn(c, e)
		return
	}
	if e := service.TaskTmpl.Update(p.Name, p.Content, p.Brief, util.User(c).Name, p.Icon, id); e != nil {
		util.OpFailFn(c, e)
		return
	}
	c.JSON(200, util.RetOK)
}

/**
 * description: 删除一个任务模板
 */
func DeleteTaskTmpls(c *gin.Context) {
	var id = cast.ToInt(c.Param("id"))
	if e := service.TaskTmpl.Delete(id); e != nil {
		util.OpFailFn(c, e)
	}
	c.JSON(200, util.RetOK)
}
