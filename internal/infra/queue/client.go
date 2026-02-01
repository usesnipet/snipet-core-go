package queue

import (
	"github.com/hibiken/asynq"
	"github.com/usesnipet/snipet-core-go/internal/config"
)

func NewAsynqClient() *asynq.Client {
	env := config.GetEnv()

	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     env.REDIS_ADDR,
		Username: env.REDIS_USER,
		Password: env.REDIS_PASS,
		DB:       1,
	})
}
