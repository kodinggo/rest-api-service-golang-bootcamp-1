package worker

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

type AsynqClient struct {
	client *asynq.Client
}

func InitAsynqClient(redistHost string) *AsynqClient {
	// 1. Connect ke redis server
	redisConn := asynq.RedisClientOpt{Addr: redistHost}

	// 2. Init asynq client (publisher)
	return &AsynqClient{
		client: asynq.NewClient(redisConn),
	}
}

func (c *AsynqClient) SendEmail(payload SendEmailPayload) (*asynq.TaskInfo, error) {
	// 3. Prepare payload
	bPayload, _ := json.Marshal(payload)

	// 4. Create task
	queueTask := asynq.NewTask(SendEmailTask, bPayload)

	return c.client.Enqueue(queueTask)
}

func (c *AsynqClient) UploadImage(payload UploadImagePayload) (*asynq.TaskInfo, error) {
	// 3. Prepare payload
	bPayload, _ := json.Marshal(payload)

	// 4. Create task
	queueTask := asynq.NewTask(UploadImageTask, bPayload)

	return c.client.Enqueue(queueTask)
}
