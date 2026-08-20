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
	// 部分清理失败：不要隐藏为成功，也不要标记为可重试——重试会重复清理。
	// 将错误上抛，由调用方人工介入恢复。
	return CleanupStatus{Retryable: false}, err
}
