package queue

import (
	"context"

	"github.com/hibiken/asynq"
	"go.uber.org/fx"
)

type Task interface {
	Name() string
	Handle(ctx context.Context, t *asynq.Task) error
}

type TasksIn struct {
	fx.In

	Tasks []Task `group:"tasks"`
}
