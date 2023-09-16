package main

import (
	"fmt"

	"root/Tasks.go"

	"github.com/hibiken/asynq"
)

func main() {
	redisCon := asynq.RedisClientOpt{
		DB:       0,
		Addr:     "127.0.0.1:6379",
		Password: "",
	}
	worker := asynq.NewServer(redisCon, asynq.Config{
		Concurrency: 12,
		Queues:      map[string]int{"user-detail": 12},
	})
	mux := asynq.NewServeMux()
	mux.HandleFunc(Tasks.TaskType, Tasks.HandleUserdtl)
	worker.Run(mux)

	fmt.Println("Main Ends! ")
}
