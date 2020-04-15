package service

import (
	"github.com/xiaomi/naftis/src/api/log"
	"github.com/xiaomi/naftis/src/api/model"
	"github.com/xiaomi/naftis/src/api/storer/db"
)

/**
 * description: taskTmpl 任务模板的增删改查
 */
var TaskTmpl taskTmpl

type taskTmpl struct{}

func (taskTmpl) Get(name, content, operator string, ids []uint, ctmin, ctmax int, revision, tp uint) []model.TaskTmpl {
	tmpls := db.GetTaskTmpl(name, content, operator, ids, ctmin, ctmax, revision, tp)
	var e error
	for i := range tmpls {
		tmpls[i].VarMap = TaskTmplVar.Get("", "", "", "", 0, tmpls[i].ID, nil)
		if e != nil {
			log.Error("[taskTmpl] Get taskTmpl fail", "err", e)
			return []model.TaskTmpl{}
		}
	}
	return tmpls
}

func (taskTmpl) Add(name, content, brief, operator, icon string, vars []model.TaskTmplVar) (model.TaskTmpl, error) {
	return db.AddTaskTmpl(name, content, brief, operator, vars, icon)
}

func (taskTmpl) Update(name, content, brief, operator, icon string, id uint) error {
	return db.UpdateTaskTmpl(name, content, brief, operator, id, icon)
}

func (taskTmpl) Delete(id int) error {
	return db.DelTaskTmpl(id)
}
