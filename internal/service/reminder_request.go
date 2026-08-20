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
	return s.client.Send(context.Background(), id)
}
