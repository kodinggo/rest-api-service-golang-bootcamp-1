package worker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const (
	SendEmailTask   = "SendEmail"
	UploadImageTask = "UploadImage"
)

type (
	SendEmailPayload struct {
		From    string
		To      string
		Subject string
		Message string
	}

	UploadImagePayload struct {
		ImageSource string
		TargetPath  string
	}
)

type taskHandler struct {
}

func (h *taskHandler) sendEmailHandler(ctx context.Context, t *asynq.Task) error {
	time.Sleep(2 * time.Second)

	var payload SendEmailPayload
	json.Unmarshal(t.Payload(), &payload)

	log.Printf("Handle send email for subject %s", payload.Subject)

	return nil
}

func (h *taskHandler) uploadImageHandler(ctx context.Context, t *asynq.Task) error {
	time.Sleep(2 * time.Second)

	var payload UploadImagePayload
	json.Unmarshal(t.Payload(), &payload)

	log.Printf("Handle upload image, target path = %s", payload.TargetPath)

	return nil
}
