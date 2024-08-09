package worker

import (
	"github.com/hibiken/asynq"
)

func InitAsynqServer(redistHost string) error {
	// 1. Connect ke redis server
	redisConn := asynq.RedisClientOpt{Addr: redistHost}

	// 2. Init asynq server
	asynqServer := asynq.NewServer(redisConn, asynq.Config{
		Concurrency: 2,
	})

	// 3. Init server mux
	var th taskHandler
	mux := asynq.NewServeMux()
	mux.HandleFunc(SendEmailTask, th.sendEmailHandler)
	mux.HandleFunc(UploadImageTask, th.uploadImageHandler)

	return asynqServer.Run(mux)
}
