package service

import (
	"errors"

	"todolist/internal/retry"
)

type DeliveryResult struct {
	Retryable bool
}

type ReminderDelivery struct {
	runner *retry.Runner
}

func NewReminderDelivery(runner *retry.Runner) *ReminderDelivery {
	return &ReminderDelivery{runner: runner}
}

func (s *ReminderDelivery) Send(id string) (DeliveryResult, error) {
	err := s.runner.Deliver(id)
	if err == nil {
		return DeliveryResult{}, nil
	}
	var committed *retry.CommitError
	if errors.As(err, &committed) && committed.Committed {
		return DeliveryResult{Retryable: false}, err
	}
	return DeliveryResult{Retryable: true}, err
}
