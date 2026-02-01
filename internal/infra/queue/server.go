package queue

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"github.com/usesnipet/snipet-core-go/internal/config"
	"go.uber.org/fx"
)

func NewAsynqServer(
	lc fx.Lifecycle,
	tasksIn TasksIn,
) *asynq.Server {

	env := config.GetEnv()

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     env.REDIS_ADDR,
			Username: env.REDIS_USER,
			Password: env.REDIS_PASS,
			DB:       1,
		},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"default": 1,
			},
		},
	)

	mux := asynq.NewServeMux()

	for _, task := range tasksIn.Tasks {
		mux.HandleFunc(task.Name(), task.Handle)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("🚀 Asynq worker started")
				if err := srv.Run(mux); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("🛑 Asynq worker stopping")
			srv.Shutdown()
			return nil
		},
	})

	return srv
}
