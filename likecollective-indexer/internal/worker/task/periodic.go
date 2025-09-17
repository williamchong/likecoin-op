package task

import (
	"errors"

	"github.com/hibiken/asynq"
)

var ErrConfigScheduler = errors.New("err config scheduler")

type Periodic interface {
	Worker
	ConfigScheduler(cronspec string, scheduler *asynq.Scheduler, options ...asynq.Option) (*asynq.Scheduler, error)
}

type periodic struct {
	taskDefs []PeriodicTaskDef
}

func DefinePeriodic(
	taskDefs []PeriodicTaskDef,
) Periodic {
	return &periodic{
		taskDefs,
	}
}

func (w *periodic) ConfigServerMux(mux *asynq.ServeMux) *asynq.ServeMux {
	for _, task := range w.taskDefs {
		mux.HandleFunc(task.GetType(), task.GetHandler())
	}
	return mux
}

func (w *periodic) ConfigScheduler(cronspec string, scheduler *asynq.Scheduler, options ...asynq.Option) (*asynq.Scheduler, error) {
	for _, taskDef := range w.taskDefs {
		task, err := taskDef.MakeTask()
		if err != nil {
			return nil, errors.Join(ErrConfigScheduler, err)
		}
		_, err = scheduler.Register(cronspec, task, options...)
		if err != nil {
			return nil, errors.Join(ErrConfigScheduler, err)
		}
	}
	return scheduler, nil
}
