package executor

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/xiaomi/naftis/src/api/log"
	"github.com/xiaomi/naftis/src/api/model"
	"github.com/xiaomi/naftis/src/api/storer/db"
	"github.com/xiaomi/naftis/src/api/util"

	"github.com/ghodss/yaml"
	"github.com/hashicorp/go-multierror"
	"istio.io/istio/pilot/pkg/config/kube/crd"
	istiomodel "istio.io/istio/pilot/pkg/model"
)

var (
	// istio 中找不到指定的task
	ErrTaskNotExists = errors.New("task isn't exists")
)

type istiocrdExecutor struct {
	client *crd.Client
}

/**
 * description: 返回一个 istiocrd executor
 */
func NewCrdExecutor() Executor {
	c, e := crd.NewClient(util.Kubeconfig(), "", istiomodel.IstioConfigTypes, "")
	if e != nil {
		log.Panic("[executor] init istiocrd fail", "error", e)
	}
	return &istiocrdExecutor{
		client: c,
	}
}

var (
	/**
	 * description: 新增task
	 */
	addTask = func(task *Task) (e error) {
		e = db.AddTask(task.TaskTmplID, task.Content, task.Operator, task.ServiceUID, task.PrevState, task.Namespace, task.Status)
		if e != nil {
			log.Error("[executor] addTask fail", "task", task, "error", e)
		}
		Push2TaskStatusCh(*task)
		return
	}
)

// 实现 Executor.Execute()，task的执行方法
func (i *istiocrdExecutor) Execute(task Task) error {
	switch task.Command {
	case model.Create, model.Replace, model.Delete, model.Rollback:
		return i.crdExec(task, addTask)
	case model.Apply:
		return i.apply(task, addTask)
	}
	return nil
}

func (i *istiocrdExecutor) create(varr []istiomodel.Config, task *Task) (errs error) {
	for _, config := range varr {
		var err error
		if config.Namespace, err = handleNamespaces(task.Namespace); err != nil {
			return err
		}

		var rev string
		if rev, err = i.client.Create(config); err != nil {
			// if the config create fail, break loop and return error
			log.Info("Created config fail", "key", config.Key(), "config", config, "error", err)
			return err
		}
		log.Info("Created config success", "key", config.Key(), "revision", rev, "config", config)
	}
	return nil
}

func (i *istiocrdExecutor) replace(varr []istiomodel.Config, task *Task) (errs error) {
	currentCfgs := make([]istiomodel.Config, 0)
	defer func() {
		task.PrevState = i.yamlOutput(currentCfgs)
	}()

	for _, config := range varr {
		var err error
		// overwrite config.Namespace with user specified namespace
		if config.Namespace, err = handleNamespaces(task.Namespace); err != nil {
			return err
		}

		// fill up revision
		if config.ResourceVersion == "" {
			current, exists := i.client.Get(config.Type, config.Name, config.Namespace)
			if !exists {
				log.Error("Task not exists", "type", config.Type, "name", config.Name, "namespace", task.Namespace)
				return ErrTaskNotExists
			}
			config.ResourceVersion = current.ResourceVersion
			// clear resourceVersion for rollback
			current.ResourceVersion = ""
			currentCfgs = append(currentCfgs, *current)
		}
		var newRev string
		if newRev, err = i.client.Update(config); err != nil {
			// if the config create fail, break loop and return error
			log.Info("Replace config fail", "key", config.Key(), "config", config, "error", err)
			return err
		}
		log.Info("Replace config success", "key", config.Key(), "config", config, "revision", newRev)
	}

	return nil
}

func (i *istiocrdExecutor) delete(varr []istiomodel.Config, task *Task) (errs error) {
	for _, config := range varr {
		var err error
		if config.Namespace, err = handleNamespaces(config.Namespace); err != nil {
			return err
		}

		if err := i.client.Delete(config.Type, config.Name, config.Namespace); err != nil {
			log.Info("Delete config fail", "key", config.Key(), "config", config, "error", err)
			// if the config delete fail, continue loop
			errs = multierror.Append(errs, fmt.Errorf("cannot delete %s: %v", config.Key(), err))
		} else {
			log.Info("Delete config success", "key", config.Key(), "config", config)
		}
	}
	return nil
}

func (i *istiocrdExecutor) crdExec(task Task, t taskDbHandler) (errs error) {
	// 将task状态标志为TaskStatusSucc
	task.Status = model.TaskStatusSucc
	// 若执行过程中发生错误，则将task状态标志为TaskStatusFail，并且执行参数taskDbHandler方法
	defer func() {
		if errs != nil {
			task.Status = model.TaskStatusFail
		}
		// 执行参数taskDbHandler方法
		t(&task)
	}()

	// ignore k8s configuration. TODO support k8s configuration
	varr, _, err := crd.ParseInputs(task.Content)
	if err != nil {
		return err
	}
	if len(varr) == 0 {
		return errors.New("nothing to execute")
	}

	switch task.Command {
	case model.Create:
		return i.create(varr, &task)
	case model.Delete:
		return i.delete(varr, &task)
	case model.Replace, model.Rollback: // NOTICE, task.Content should be prevState
		return i.replace(varr, &task)
	}

	return nil
}

var (
	namespace        string
	defaultNamespace = "default"
)

func handleNamespaces(objectNamespace string) (string, error) {
	if objectNamespace != "" && namespace != "" && namespace != objectNamespace {
		return "", fmt.Errorf(`the namespace from the provided object "%s" does `+
			`not match the namespace "%s". You must pass '--namespace=%s' to perform `+
			`this operation`, objectNamespace, namespace, objectNamespace)
	}

	if namespace != "" {
		return namespace, nil
	}

	if objectNamespace != "" {
		return objectNamespace, nil
	}
	return defaultNamespace, nil
}

func (i *istiocrdExecutor) apply(task Task, t taskDbHandler) (errs error) {
	task.Status = model.TaskStatusSucc
	defer func() {
		if errs != nil {
			task.Status = model.TaskStatusFail
		}
		t(&task)
	}()

	// ignore k8s configuration. TODO support k8s configuration
	varr, _, err := crd.ParseInputs(task.Content)
	if err != nil {
		return err
	}
	if len(varr) == 0 {
		return errors.New("nothing to execute")
	}

	if err := i.create(varr, &task); err != nil {
		return i.replace(varr, &task)
	}

	return
}

func (i *istiocrdExecutor) yamlOutput(configList []istiomodel.Config) string {
	buf := bytes.NewBuffer([]byte{})
	descriptor := i.client.ConfigDescriptor()
	for _, config := range configList {
		schema, exists := descriptor.GetByType(config.Type)
		if !exists {
			fmt.Printf("Unknown kind %q for %v", crd.ResourceName(config.Type), config.Name)
			continue
		}
		obj, err := crd.ConvertConfig(schema, config)
		if err != nil {
			fmt.Printf("Could not decode %v: %v", config.Name, err)
			continue
		}
		bytes, err := yaml.Marshal(obj)
		if err != nil {
			fmt.Printf("Could not convert %v to YAML: %v", config, err)
			continue
		}

		buf.Write(bytes)
		buf.WriteString("---")
	}

	return buf.String()
}
