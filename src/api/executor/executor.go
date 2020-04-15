package executor

import (
	"errors"
	"sync"

	"github.com/xiaomi/naftis/src/api/model"
)

// 默认task执行器
var DefaultExecutor Executor

/**
 * description: 初始化executor包
 */
func Init() {
	// 返回一个 istiocrd executor
	DefaultExecutor = NewCrdExecutor()
}

// Executor 将从通道分派并执行任务
type Executor interface {
	Execute(Task) error
}

// Task is an alias of model.Task
type Task = model.Task

type taskDbHandler = func(task *Task) error

var (
	// invalid command error
	ErrUnknownCmd = errors.New("unknown command")
)

var (
	// TaskStatusChM 将task执行结果保存到 channel map 中
	TaskStatusChM    = make(map[string]chan Task)
	taskStatusChMMtx = new(sync.RWMutex)
)

/**
 * description: 返回 task channel from taskStatusChM group by operator, 如果channel不存在，则新建一个
 */
func GetOrAddTaskStatusChM(name string) chan Task {
	taskStatusChMMtx.Lock()
	u, ok := TaskStatusChM[name]
	if !ok {
		u = make(chan Task, 1000)
		TaskStatusChM[name] = u
	}
	taskStatusChMMtx.Unlock()
	return u
}

/**
 * description: 将 task 添加到 taskStatusChM 中
 */
func Push2TaskStatusCh(task Task) {
	GetOrAddTaskStatusChM(task.Operator) <- task
}
