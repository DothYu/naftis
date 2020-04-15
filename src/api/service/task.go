package service

import (
	"github.com/xiaomi/naftis/src/api/model"
	"github.com/xiaomi/naftis/src/api/storer/db"
	"github.com/xiaomi/naftis/src/api/worker"
)

/**
 * description: Task的增删改查
 */
var Task task

type task struct{}

func (task) Get(name, content, operator, serviceUID string, id uint, ctmin, ctmax int, revision uint) []model.Task {
	return db.GetTask(name, content, operator, serviceUID, id, ctmin, ctmax, revision)
}

func (task) Add(tmplID uint, command int, content, operator, serviceUID, namespace string) error {
	return worker.Feed(tmplID, command, content, operator, serviceUID, namespace, 1)
}

// Deprecated: the function is already Deprecated.
func (task) Update(content, operator, serviceUID string, id, tmplID uint) error {
	return db.UpdateTask(content, operator, serviceUID, id, tmplID, 0)
}

// Deprecated: the function is already Deprecated.
func (task) Delete(id uint, operator string) error {
	return db.DeleteTask(id, operator)
}
