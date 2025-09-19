package task

import (
	"context"
	"fmt"
	"maps"

	"github.com/hibiken/asynq"
)

type AsynqHandler func(context.Context, *asynq.Task) error

type TaskDef interface {
	GetType() string
	GetHandler() AsynqHandler
}

type taskDef struct {
	taskType string
	handler  AsynqHandler
}

func DefineTask(
	taskType string,
	handler AsynqHandler,
) TaskDef {
	return &taskDef{
		taskType,
		handler,
	}
}

func (t *taskDef) GetType() string {
	return t.taskType
}

func (t *taskDef) GetHandler() AsynqHandler {
	return t.handler
}

type TaskDefsByName map[string]TaskDef

func (m TaskDefsByName) Register(def TaskDef) TaskDef {
	_, ok := m[def.GetType()]
	if ok {
		panic(fmt.Errorf("err register task: task %s already registered", def.GetType()))
	}
	m[def.GetType()] = def
	return def
}

func (m TaskDefsByName) GetRegisteredTasks() []string {
	res := make([]string, 0)
	for k := range maps.Keys(m) {
		res = append(res, k)
	}
	return res
}

func (m TaskDefsByName) MakeWorker(types ...string) (Worker, error) {
	var taskDefs = make([]TaskDef, len(types))
	for i, t := range types {
		if taskDef, ok := m[t]; ok {
			taskDefs[i] = taskDef
		} else {
			return nil, fmt.Errorf("err unknown type %s", t)
		}
	}
	return DefineWorker(taskDefs), nil
}
