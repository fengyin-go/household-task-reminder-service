package service

import (
	"context"

	"todolist/internal/dispatch"
)

type ReminderScheduler struct {
	dispatcher *dispatch.Service
}

func NewReminderScheduler(dispatcher *dispatch.Service) *ReminderScheduler {
	return &ReminderScheduler{dispatcher: dispatcher}
}

func (s *ReminderScheduler) Schedule(ctx context.Context, job dispatch.Job) error {
	return s.dispatcher.Schedule(context.Background(), job)
}
