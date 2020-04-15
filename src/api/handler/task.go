package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/xiaomi/naftis/src/api/model"
	"github.com/xiaomi/naftis/src/api/service"
	"github.com/xiaomi/naftis/src/api/util"
)

type taskPayload struct {
	ID         uint     `json:"id"`
	TmplID     uint     `json:"tmplID"`
	Command    string   `json:"command"`
	Content    string   `json:"content"`
	ServiceUID string   `json:"serviceUID"`
	VarMaps    []string `json:"varMaps"`
	Namespace  string   `json:"namespace"`
}

var (
	// invalid service UID
	ErrInvalidServiceUID = errors.New("invalid serviceUID")
	// invalid variable maps
	ErrInvalidVarMap = errors.New("invalid varMap")
	// invalid task template ID
	ErrInvalidTmplID = errors.New("invalid tmplID")
	// invalid command
	ErrInvalidCommand = errors.New("invalid command")
	// invalid namespace
	ErrInvalidNamespace = errors.New("invalid namespace")
)

/**
 * description: taskPayload验证
 */
func (t taskPayload) validate() (e error) {
	if t.ServiceUID == "" {
		return ErrInvalidServiceUID
	}
	if t.Namespace == "" {
		return ErrInvalidNamespace
	}
	if t.Command == "" {
		return ErrInvalidCommand
	}
	return
}

/**
 * description: 根据条件返回已保存的tasks，及其所属task模板(if exist)
 */
func ListTasks(c *gin.Context) {
	var id = cast.ToUint(c.Param("id"))
	var serviceUID = c.Query("serviceUID")

	// query tasks and get related tasktmpl ids
	tasks := service.Task.Get("", "", "", serviceUID, id, 0, 0, 0)
	tmplIDs := make([]uint, 0, len(tasks))
	for _, t := range tasks {
		tmplIDs = append(tmplIDs, t.TaskTmplID)
	}

	// query tasktmpls by tasktmpl ids, then join tasktmpl to task by tasktmpl id
	tmpls := service.TaskTmpl.Get("", "", "", tmplIDs, 0, 0, 0, 0)
	type item struct {
		model.Task
		TaskTmpl model.TaskTmpl `json:"tmpl"`
	}
	var data = make([]item, 0, len(tasks))
	for _, t := range tasks {
		var hasTpl bool
		for _, tmpl := range tmpls {
			if t.TaskTmplID == tmpl.ID {
				data = append(data, item{
					Task:     t,
					TaskTmpl: tmpl,
				})
				hasTpl = true
				break
			}
		}
		if !hasTpl {
			data = append(data, item{
				Task: t,
			})
		}
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

func convertCmd(cmdStr string) (cmd int) {
	switch cmdStr {
	case "apply":
		cmd = int(model.Apply)
	case "delete":
		cmd = int(model.Delete)
	case "rollback":
		cmd = int(model.Rollback)
	}
	return
}

// AddTasks adds a task into task worker.
/**
 * description: 添加一个任务到任务工作器中
 */
func AddTasks(c *gin.Context) {
	var p taskPayload
	if e := c.BindJSON(&p); e != nil {
		util.BindFailFn(c, e)
		return
	}

	// validate payload
	if e := p.validate(); e != nil {
		util.BindFailFn(c, e)
		return
	}

	// 执行rollback回滚操作
	if cmd := convertCmd(p.Command); cmd == int(model.Rollback) {
		if e := service.Task.Add(0, cmd, p.Content, util.User(c).Name, p.ServiceUID, p.Namespace); e != nil {
			util.OpFailFn(c, e)
			return
		}
		return
	}

	// Execute apply command.
	if len(p.VarMaps) == 0 {
		util.BindFailFn(c, ErrInvalidVarMap)
		return
	}
	if p.TmplID == 0 {
		util.BindFailFn(c, ErrInvalidTmplID)
		return
	}
	// get task template by template id.
	tmplIDs := make([]uint, 0, 1)
	if p.TmplID != 0 {
		tmplIDs = append(tmplIDs, p.TmplID)
	}
	tmpls := service.TaskTmpl.Get("", "", "", tmplIDs, 0, 0, 0, 0)
	if len(tmpls) == 0 {
		util.BindFailFn(c, errors.New("invalid tmplID"))
		return
	}
	var content string
	for _, m := range p.VarMaps {
		ct, e := model.ExecTmpl(tmpls[0], m)
		if e != nil {
			util.OpFailFn(c, e)
			return
		}
		content += ct + "---\n"
	}

	// feed task to worker
	if e := service.Task.Add(p.TmplID, convertCmd(p.Command), content, util.User(c).Name, p.ServiceUID, p.Namespace); e != nil {
		util.OpFailFn(c, e)
		return
	}

	c.JSON(200, util.RetOK)
}

/**
 * description: 更新指定的task
 */
// Deprecated: the function is already Deprecated.
func UpdateTasks(c *gin.Context) {
	var id = cast.ToUint(c.Param("id"))
	var p taskPayload
	if e := c.BindJSON(&p); e != nil {
		util.BindFailFn(c, e)
		return
	}
	if e := service.Task.Update(p.Content, util.User(c).Name, p.ServiceUID, id, p.TmplID); e != nil {
		util.OpFailFn(c, e)
	}

	c.JSON(200, util.RetOK)
}

/**
 * description: 删除指定的task
 */
// Deprecated: the function is already Deprecated.
func DeleteTasks(c *gin.Context) {
	var id = cast.ToUint(c.Param("id"))
	if e := service.Task.Delete(id, util.User(c).Name); e != nil {
		util.OpFailFn(c, e)
	}
	c.JSON(200, util.RetOK)
}
