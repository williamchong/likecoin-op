package task

import (
	"github.com/hibiken/asynq"
)

type Worker interface {
	ConfigServerMux(*asynq.ServeMux) *asynq.ServeMux
}

type worker struct {
	tasks []TaskDef
}

func DefineWorker(
	tasks []TaskDef,
) Worker {
	return &worker{
		tasks,
	}
}

func (w *worker) ConfigServerMux(mux *asynq.ServeMux) *asynq.ServeMux {
	for _, task := range w.tasks {
		mux.HandleFunc(task.GetType(), task.GetHandler())
	}
	return mux
}
