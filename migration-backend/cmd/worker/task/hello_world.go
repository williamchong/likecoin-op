package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

const TypeHelloWorld = "hello_world"

type HelloWorldPayload struct {
	Message string
}

func NewHelloWorldTask(message string) (*asynq.Task, error) {
	payload, err := json.Marshal(HelloWorldPayload{Message: message})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeHelloWorld, payload), nil
}

func HandleHelloWorldTask(ctx context.Context, t *asynq.Task) error {
	var p HelloWorldPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Hello %v", p.Message)
	return nil
}
