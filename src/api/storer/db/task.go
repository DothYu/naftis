package db

import (
	"github.com/xiaomi/naftis/src/api/log"
	"github.com/xiaomi/naftis/src/api/model"

	"github.com/jinzhu/gorm"
)

/**
 * description: 【新增】往`tasks`中新增一条记录，即新增一个task
 */
func AddTask(tmplID uint, content, operator, serviceUID, prevState, namespace string, status uint) error {
	if content == "" || operator == "" || serviceUID == "" || namespace == "" {
		return ErrInvalidParams
	}

	// insert record into `tasks`
	var task = model.Task{
		TaskTmplID: tmplID,
		Content:    content,
		Operator:   operator,
		Revision:   1,
		Status:     status,
		ServiceUID: serviceUID,
		PrevState:  prevState,
		Namespace:  namespace,
	}
	if e := db.Create(&task).Error; e != nil {
		log.Error("[service] AddTask fail", "error", e.Error(), "record", task)
	}

	return nil
}

/**
 * description: 【删除】删除`tasks`中一条记录，即删除一个指定的task
 */
// Deprecated: the function is already Deprecated.
func DeleteTask(id uint, operator string) error {
	if e := db.Where("id = ?", id).Delete(model.Task{}).Update("operator", operator).Error; e != nil {
		log.Info("[service] DeleteTask fail", "error", e.Error())
		return e
	}
	return nil
}

/**
 * description: 【更新】更新`tasks`中一条记录，即更新一个指定的task
 */
// Deprecated: the function is already Deprecated.
func UpdateTask(content, operator, serviceUID string, id, tmplID, status uint) error {
	if id == 0 {
		return ErrInvalidParams
	}

	udpates := map[string]interface{}{}
	if content != "" {
		udpates["content"] = content
	}
	if operator != "" {
		udpates["operator"] = operator
	}
	if serviceUID != "" {
		udpates["service_uid"] = serviceUID
	}
	if tmplID != 0 {
		udpates["task_tmpl_id"] = tmplID
	}
	if status != 0 {
		udpates["status"] = status
	}
	udpates["revision"] = gorm.Expr("revision + 1")

	// update record of `tasks`
	var t = model.Task{}
	if e := db.Model(&t).Where("id = ?", id).Updates(udpates).Error; e != nil {
		log.Info("[service] UpdateTask fail", "error", e.Error(), "record", udpates)
	}

	return nil
}

/**
 * description: 【查询】根据提供的字段查询`tasks`中的记录
 */
func GetTask(name, content, operator, serviceUID string, id uint, ctmin, ctmax int, revision uint) []model.Task {
	var whereStr = "1=1 "
	var args = make([]interface{}, 0)
	var tasks = make([]model.Task, 0)

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
	if serviceUID != "" {
		whereStr += "and service_uid like ?"
		args = append(args, serviceUID)
	}
	if id != 0 {
		whereStr += "and id = ?"
		args = append(args, id)
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

	// 通过gorm操作数据库查询
	if e := db.Where(whereStr, args...).Order("created_at desc").Find(&tasks).Error; e != nil {
		log.Info("[service] GetTask fail", "error", e.Error())
	}

	return tasks
}
