package service

import (
	"context"

	"todolist/internal/dispatch"
)

// ReminderScheduler is the application-facing boundary for dispatching a reminder.
type ReminderScheduler struct {
	dispatcher *dispatch.Service
}

func NewReminderScheduler(dispatcher *dispatch.Service) *ReminderScheduler {
	return &ReminderScheduler{dispatcher: dispatcher}
}

func (s *ReminderScheduler) Schedule(ctx context.Context, job dispatch.Job) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return s.dispatcher.Schedule(ctx, job)
}
