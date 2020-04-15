package service

import (
	"strings"

	"github.com/xiaomi/naftis/src/api/bootstrap"
	"github.com/xiaomi/naftis/src/api/model"
	"github.com/xiaomi/naftis/src/api/storer/db"
)

/**
 * description: taskTmplVar的增删改查
 */
var TaskTmplVar taskTmplVar

type taskTmplVar struct{}

var mockVersions = []string{
	"v1", "v2", "v3",
}

func (taskTmplVar) Get(name, title, comment, dataSource string, formType, tasktmplID uint, ids []uint) []model.TaskTmplVar {
	vars := db.GetTaskTmplVar(name, title, comment, dataSource, formType, tasktmplID, ids)
	for i := range vars {
		// 解析`datasource`变量，返回特定的`datasource`映射
		switch strings.ToLower(vars[i].DataSource) {
		case "host":
			data := make(map[string]string)
			svcs := ServiceInfo.Services("").Exclude("kube-system", bootstrap.Args.IstioNamespace, bootstrap.Args.Namespace)
			for _, s := range svcs {
				data[s.Name] = s.Name
			}
			vars[i].Data = data
		case "namespace":
			data := make(map[string]string)
			ns := ServiceInfo.Namespaces("").Exclude("kube-system", bootstrap.Args.IstioNamespace, bootstrap.Args.Namespace)
			for _, s := range ns {
				data[s.Name] = s.Name
			}
			vars[i].Data = data
		case "version":
			vars[i].Data = mockVersions
		}
	}
	return vars
}

func (taskTmplVar) Add(name, title, comment, dataSource string, taskTmplID, formType uint) (model.TaskTmplVar, error) {
	return db.AddTaskTmplVar(name, title, comment, dataSource, taskTmplID, formType)
}

func (taskTmplVar) Update(name, title, comment, dataSource string, id, formType uint) error {
	return db.UpdateTaskTmplVar(name, title, comment, dataSource, id, formType)
}

func (taskTmplVar) Delete(id int) error {
	return db.DelTaskTmplVar(id)
}
