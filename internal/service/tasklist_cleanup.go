package service

import "todolist/internal/cleanup"

type CleanupStatus struct{ Retryable bool }

type TaskListCleanup struct{ retry *cleanup.Retry }

func NewTaskListCleanup(retry *cleanup.Retry, audit ...any) *TaskListCleanup { return &TaskListCleanup{retry: retry} }

func (s *TaskListCleanup) Delete(id string) (CleanupStatus, error) {
	err := s.retry.Run(id)
	if err == nil {
		return CleanupStatus{}, nil
	}
	return CleanupStatus{Retryable: true}, err
}
