package service

import (
	"errors"

	"todolist/internal/cleanup"
)

type CleanupStatus struct{ Retryable bool }

type TaskListCleanup struct {
	retry *cleanup.Retry
	audit *cleanup.Audit
}

func NewTaskListCleanup(retry *cleanup.Retry, audits ...*cleanup.Audit) *TaskListCleanup {
	var audit *cleanup.Audit
	if len(audits) > 0 {
		audit = audits[0]
	}
	if audit == nil {
		audit = &cleanup.Audit{}
	}
	return &TaskListCleanup{retry: retry, audit: audit}
}

func (s *TaskListCleanup) Delete(id string) (CleanupStatus, error) {
	err := s.retry.Run(id)
	s.audit.Record(err)
	var partial *cleanup.PartialError
	if errors.As(err, &partial) && partial.Cleaned {
		return CleanupStatus{}, err
	}
	return CleanupStatus{Retryable: err != nil}, err
}
