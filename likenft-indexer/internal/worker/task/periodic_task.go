package task

import (
	"fmt"
	"maps"

	"github.com/hibiken/asynq"
)

type TaskGenerator func() (*asynq.Task, error)

type PeriodicTaskDef interface {
	TaskDef
	MakeTask() (*asynq.Task, error)
}

type periodicTaskDef struct {
	task          TaskDef
	taskGenerator TaskGenerator
}

func DefinePeriodicTask(
	task TaskDef,
	taskGenerator TaskGenerator,
) PeriodicTaskDef {
	return &periodicTaskDef{
		task,
		taskGenerator,
	}
}

func (t *periodicTaskDef) GetType() string {
	return t.task.GetType()
}

func (t *periodicTaskDef) GetHandler() AsynqHandler {
	return t.task.GetHandler()
}

func (t *periodicTaskDef) MakeTask() (*asynq.Task, error) {
	return t.taskGenerator()
}

type PeriodicTaskByName map[string]PeriodicTaskDef

func (m PeriodicTaskByName) Register(def PeriodicTaskDef) PeriodicTaskDef {
	_, ok := m[def.GetType()]
	if ok {
		panic(fmt.Errorf("err register periodic task: task %s already registered", def.GetType()))
	}
	m[def.GetType()] = def
	return def
}

func (m PeriodicTaskByName) GetRegisteredPeriodicTasks() []string {
	res := make([]string, 0)
	for k := range maps.Keys(m) {
		res = append(res, k)
	}
	return res
}

func (m PeriodicTaskByName) MakePeriodic(types ...string) (Periodic, error) {
	var taskDefs = make([]PeriodicTaskDef, len(types))
	for i, t := range types {
		if taskDef, ok := m[t]; ok {
			taskDefs[i] = taskDef
		} else {
			return nil, fmt.Errorf("err unknown type %s", t)
		}
	}
	return DefinePeriodic(taskDefs), nil
}
