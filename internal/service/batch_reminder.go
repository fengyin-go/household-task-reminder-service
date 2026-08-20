package service

import "todolist/internal/batch"

type BatchReminderService struct{ coordinator batch.Coordinator }

func (s BatchReminderService) Dispatch(items []batch.Item) ([]batch.Result, error) {
	if len(items) == 0 {
		return []batch.Result{}, nil
	}
	return s.coordinator.Run(items)
}
