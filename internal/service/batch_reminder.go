package service

import "todolist/internal/batch"

type BatchReminderService struct{ coordinator batch.Coordinator }

func (s BatchReminderService) Dispatch(items []batch.Item) ([]batch.Result, error) {
	return s.coordinator.Run(items)
}
