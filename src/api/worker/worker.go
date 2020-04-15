package worker

import (
	"fmt"

	"github.com/xiaomi/naftis/src/api/executor"
	"github.com/xiaomi/naftis/src/api/log"
	"github.com/xiaomi/naftis/src/api/model"
)

type worker struct {
	jobs  chan executor.Task
	stop  func()
	done  chan bool
	block chan bool
}

const (
	// 定义任务队列大小
	JobQueueSize = 1000
)

var w = &worker{
	jobs:  make(chan executor.Task, JobQueueSize),
	block: make(chan bool, 1),
}

// 开启任务执行器
func Start() {
	for {
		select {
		case job := <-w.jobs:
			if job.Command != 0 {
				if e := executor.DefaultExecutor.Execute(job); e != nil {
					log.Error("[worker] execute fail", "error", e)
				}
			}
		}
	}
}

/**
 * description: 停止worker并关闭任务队列
 */
func Stop() {
	fmt.Println(`terminating worker`)
	w.block <- true
	close(w.jobs)
}

// 将一个新任务添加到任务队列中
func Feed(tmplID uint, command int, content string, operator string, serviceUID, namespace string, revision uint) error {
	select {
	case <-w.block:
		fmt.Println(`worker is terminating, can't add job any more.'`)
		w.block <- true
	default:
		t := executor.Task{
			TaskTmplID: tmplID,
			Content:    content,
			Operator:   operator,
			ServiceUID: serviceUID,
			Revision:   revision,
			Command:    model.TaskCmd(command),
			Namespace:  namespace,
		}
		// push task with processing status into TaskStatusCh
		t.Status = model.TaskStatusProcessing
		executor.Push2TaskStatusCh(t)

		// push task into jobs channel
		t.Status = model.TaskStatusDefault
		w.jobs <- t
	}
	return nil
}
