package db

import (
	"github.com/xiaomi/naftis/src/api/log"
	"github.com/xiaomi/naftis/src/api/model"

	"github.com/jinzhu/gorm"
)

/**
 * description: 新增`task_tmpl_vars`一条记录
 */
func AddTaskTmplVar(name, title, comment, dataSource string, taskTmplID, formType uint) (t model.TaskTmplVar, e error) {
	if name == "" || title == "" {
		return t, ErrInvalidParams
	}

	t = model.TaskTmplVar{
		TaskTmplID: taskTmplID,
		Name:       name,
		Title:      title,
		Comment:    comment,
		DataSource: dataSource,
		FormType:   formType,
	}

	if e := db.Create(&t).Error; e != nil {
		log.Error("[service] AddTaskTmpl fail", "error", e.Error())
	}

	return
}

/**
 * description: 删除`task_tmpl_vars`中特定的记录
 */
func DelTaskTmplVar(id int) error {
	if e := db.Where("id = ?", id).Delete(model.TaskTmplVar{}).Error; e != nil {
		log.Info("[service] TaskTmplVar fail", "error", e.Error())
		return e
	}
	return nil
}

/**
 * description: 更新`task_tmpl_vars`中特定的记录
 */
func UpdateTaskTmplVar(name, title, comment, dataSource string, id, formType uint) error {
	if id == 0 || name == "" || title == "" {
		return ErrInvalidParams
	}

	udpates := map[string]interface{}{}
	if name != "" {
		udpates["name"] = name
	}
	if title != "" {
		udpates["title"] = title
	}
	if comment != "" {
		udpates["comment"] = comment
	}
	if dataSource != "" {
		udpates["data_source"] = dataSource
	}
	if formType != 0 {
		udpates["form_type"] = formType
	}
	udpates["revision"] = gorm.Expr("revision + 1")

	if e := db.Model(model.TaskTmplVar{}).Where("id = ?", id).Update(udpates).Error; e != nil {
		log.Info("[service] UpdateTask fail", "error", e.Error())
	}

	return nil
}

/**
 * description: 根据条件查询`task_tmpl_vars`中符合的记录
 */
func GetTaskTmplVar(name, title, comment, dataSource string, formType, tasktmplID uint, ids []uint) []model.TaskTmplVar {
	var whereStr = "1=1 "
	var args = make([]interface{}, 0)
	var vars = make([]model.TaskTmplVar, 0)

	if name != "" {
		whereStr += "and name like ?"
		args = append(args, name)
	}
	if title != "" {
		whereStr += "and title like ?"
		args = append(args, title)
	}
	if comment != "" {
		whereStr += "and comment like ?"
		args = append(args, comment)
	}
	if dataSource != "" {
		whereStr += "and data_source like ?"
		args = append(args, dataSource)
	}
	if len(ids) != 0 {
		whereStr += "and id in (?)"
		args = append(args, ids)
	}
	if formType != uint(0) {
		whereStr += "and form_type = ?"
		args = append(args, formType)
	}
	if tasktmplID != uint(0) {
		whereStr += "and task_tmpl_id = ?"
		args = append(args, tasktmplID)
	}

	if e := db.Where(whereStr, args...).Find(&vars).Error; e != nil {
		log.Info("[service] GetTaskTmpl fail", "error", e.Error())
	}

	return vars
}
