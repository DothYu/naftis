package model

import (
	"time"
)

/**
 * description: 定义了taskTmpl的结构体，对应table `task_tmpls`
 */
type TaskTmpl struct {
	ID        uint          `json:"id" gorm:"primary_key"`
	CreatedAt time.Time     `json:"createAt"`
	UpdatedAt time.Time     `json:"updateAt"`
	DeletedAt *time.Time    `json:"deleteAt" sql:"index"`
	Name      string        `json:"name"`
	Content   string        `json:"content"`
	Brief     string        `json:"brief"`
	Revision  uint          `json:"revision"`
	Operator  string        `json:"operator"`
	Icon      string        `json:"icon"`
	VarMap    []TaskTmplVar `json:"varMap" gorm:"-"`
}

// Var defines template variable fields.
type Var struct {
	Name       string      `json:"name"`
	Title      string      `json:"comment" gorm:"column:title"`
	Type       string      `json:"type" gorm:"column:form_type"`
	DataSource interface{} `json:"value" gorm:"column:data_source"`
}

/**
 * description: 定义了模板变量结构体
 */
type TaskTmplVar struct {
	TaskTmplID uint        `json:"taskTmplID" gorm:"task_tmpl_id"`
	Name       string      `json:"name"`
	Title      string      `json:"title"`
	Comment    string      `json:"comment"`
	FormType   uint        `json:"formType"`
	DataSource string      `json:"dataSource" gorm:"column:data_source"`
	Default    string      `json:"default"`
	Data       interface{} `json:"data" gorm:"-"`
}

const (
	// task is currently created.
	TaskStatusDefault uint = iota
	// task 正在执行
	TaskStatusProcessing
	// task 执行成功
	TaskStatusSucc
	// task 执行失败
	TaskStatusFail
)

// TaskCmd defines istioctl subcommand.
type TaskCmd int

const (
	// Apply包含 Create和Replace命令
	// 它会首先尝试Replace命令，如果出错，则执行Create命令
	Apply TaskCmd = iota + 1
	// Create represents istioctl "create" subcommand.
	Create
	// Replace represents istioctl "create" subcommand.
	Replace
	// Delete represents istioctl "delete" subcommand.
	Delete
	// Rollback rollbacks task with prev istio resource yaml.
	Rollback
)

/**
 * description: 定义了task的结构体，对应table `tasks`
 */
type Task struct {
	ID         uint       `json:"id" gorm:"primary_key"`
	CreatedAt  time.Time  `json:"createAt"`
	UpdatedAt  time.Time  `json:"updateAt"`
	DeletedAt  *time.Time `json:"deleteAt" sql:"index"`
	Content    string     `json:"content"`
	Revision   uint       `json:"revision"`
	Operator   string     `json:"operator"`
	TaskTmplID uint       `json:"taskTmplID"`
	ServiceUID string     `json:"serviceUID"`
	Status     uint       `json:"status"`
	TaskTmpl   TaskTmpl   `gorm:"ForeignKey:TaskTmplID"`
	Command    TaskCmd    `json:"command"`
	PrevState  string     `json:"prevState"`
	Namespace  string     `json:"namespace"`
}
