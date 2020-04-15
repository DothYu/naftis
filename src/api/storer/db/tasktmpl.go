package db

import (
	"fmt"
	"strings"

	"github.com/xiaomi/naftis/src/api/log"
	"github.com/xiaomi/naftis/src/api/model"

	"github.com/jinzhu/gorm"
)

/**
 * description: `task_tmpls` 新增一个task模板
 */
func AddTaskTmpl(name, content, brief, operator string, vars []model.TaskTmplVar, icon string) (t model.TaskTmpl, e error) {
	// 开启事务
	tx := db.Begin()

	if name == "" || content == "" {
		return t, ErrInvalidParams
	}

	t = model.TaskTmpl{
		Name:     name,
		Brief:    brief,
		Content:  content + "\n",
		Revision: 1,
		Operator: operator,
		Icon:     icon,
	}

	if e := tx.Create(&t).Error; e != nil {
		log.Error("[service] AddTaskTmpl fail", "error", e.Error())
		tx.Rollback()
		return t, ErrSQLExec
	}

	for i := range vars {
		vars[i].Data = ""
	}

	valueStrings := make([]string, 0, len(vars))
	valueArgs := make([]interface{}, 0, len(vars)*3)
	for _, v := range vars {
		valueStrings = append(valueStrings, "(?,?,?,?,?,?,?)")
		valueArgs = append(valueArgs, t.ID, v.Name, v.Title, v.Comment, v.FormType, v.DataSource, v.Default)
	}

	stmt := fmt.Sprintf("INSERT INTO task_tmpl_vars (`task_tmpl_id`, `name`, `title`, `comment`, `form_type`, `data_source`, `default`) VALUES %s", strings.Join(valueStrings, ","))
	// 插入数据
	if e := tx.Exec(stmt, valueArgs...).Error; e != nil {
		log.Error("[service] AddTaskTmplVar fail", "error", e.Error())
		tx.Rollback()
		return t, ErrSQLExec
	}

	return t, tx.Commit().Error
}

/**
 * description: `task_tmpls` 删除指定task模板
 */
func DelTaskTmpl(id int) error {
	if e := db.Where("id = ?", id).Delete(model.TaskTmpl{}).Error; e != nil {
		log.Info("[service] DelTaskTmpl fail", "error", e.Error())
		return e
	}
	return nil
}

/**
 * description: `task_tmpls` 更新指定task模板
 */
func UpdateTaskTmpl(name, content, brief, operator string, id uint, icon string) error {
	if id == 0 {
		return ErrInvalidParams
	}

	udpates := map[string]interface{}{}
	if name != "" {
		udpates["name"] = name
	}
	if content != "" {
		udpates["content"] = content
	}
	if brief != "" {
		udpates["brief"] = brief
	}
	if operator != "" {
		udpates["operator"] = operator
	}
	if icon != "" {
		udpates["icon"] = icon
	}
	udpates["revision"] = gorm.Expr("revision + 1")

	if e := db.Model(model.TaskTmpl{}).Where("id = ?", id).Update(udpates).Error; e != nil {
		log.Info("[service] UpdateTask fail", "error", e.Error())
	}

	return nil
}

/**
 * description: 【查询】根据提供的条件查询`task_tmpls`中符合的记录
 */
func GetTaskTmpl(name, content, operator string, ids []uint, ctmin, ctmax int, revision, tp uint) []model.TaskTmpl {
	var whereStr = "1=1 "
	var args = make([]interface{}, 0)
	var tasktmpls = make([]model.TaskTmpl, 0)

	if name != "" {
		whereStr += "and name like ?"
		args = append(args, name)
	}
	if content != "" {
		whereStr += "and content like ?"
		args = append(args, content)
	}
	if operator != "" {
		whereStr += "and operator like ?"
		args = append(args, operator)
	}
	if len(ids) != 0 {
		whereStr += "and id in (?)"
		args = append(args, ids)
	}
	if ctmin != 0 {
		whereStr += "and create_time > ?"
		args = append(args, ctmin)
	}
	if ctmax != 0 {
		whereStr += "and create_time < ?"
		args = append(args, ctmax)
	}
	if revision != 0 {
		whereStr += "and revision = ?"
		args = append(args, revision)
	}
	if tp != 0 {
		whereStr += "and type = ?"
		args = append(args, tp)
	}

	if e := db.Where(whereStr, args...).Find(&tasktmpls).Error; e != nil {
		log.Info("[service] GetTaskTmpl fail", "error", e.Error())
	}

	return tasktmpls
}
