package service

import (
	"context"

	"todolist/internal/request"
)

type ReminderRequestService struct{ client *request.Client }

func NewReminderRequestService(client *request.Client) *ReminderRequestService {
	return &ReminderRequestService{client: client}
}

func (s *ReminderRequestService) Create(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if id == "" {
		return context.Canceled
	}
	return s.client.Send(ctx, id)
}
