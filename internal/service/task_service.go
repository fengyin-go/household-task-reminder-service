package service

import (
	"sort"
	"time"

	"todolist/internal/model"
	"todolist/pkg/idgen"
)

func (s *Service) CreateTask(t model.Task) (*model.Task, error) {
	if err := t.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetTaskList(t.TaskListID); err != nil {
		return nil, model.NewValidationError("task_list_id", "所属清单不存在")
	}
	t.ID = idgen.Hex()
	t.CreatedAt = time.Now()
	t.UpdatedAt = t.CreatedAt
	if err := s.store.CreateTask(&t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Service) GetTask(id string) (*model.Task, error) {
	return s.store.GetTask(id)
}

func (s *Service) ListTasks(filter model.TaskFilter, page, size int) ([]*model.Task, int, error) {
	all := s.store.ListTasks()
	matched := make([]*model.Task, 0, len(all))
	for _, t := range all {
		if filter.Match(t) {
			matched = append(matched, t)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Task{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateTask(id string, t model.Task) (*model.Task, error) {
	existing, err := s.store.GetTask(id)
	if err != nil {
		return nil, err
	}
	if t.Title != "" {
		existing.Title = t.Title
	}
	if t.Description != "" {
		existing.Description = t.Description
	}
	if t.TaskListID != "" {
		if _, err := s.store.GetTaskList(t.TaskListID); err != nil {
			return nil, model.NewValidationError("task_list_id", "所属清单不存在")
		}
		existing.TaskListID = t.TaskListID
	}
	if t.Priority >= 0 {
		existing.Priority = t.Priority
	}
	if !t.DueDate.IsZero() {
		existing.DueDate = t.DueDate
	}
	if err := existing.Validate(); err != nil {
		return nil, err
	}
	existing.UpdatedAt = time.Now()
	if err := s.store.UpdateTask(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *Service) DeleteTask(id string) error {
	return s.store.DeleteTask(id)
}

// TransitionTaskStatus 流转任务状态。
func (s *Service) TransitionTaskStatus(id string, toStatus string) (*model.Task, error) {
	t, err := s.store.GetTask(id)
	if err != nil {
		return nil, err
	}
	if !model.CanTransitionTask(t.Status, toStatus) {
		return nil, model.NewValidationError("status", "状态流转不合法")
	}
	t.Status = toStatus
	t.UpdatedAt = time.Now()
	if err := s.store.UpdateTask(t); err != nil {
		return nil, err
	}
	return t, nil
}

// TaskListStats 单个清单的完成率统计。
type TaskListStats struct {
	TaskListID  string `json:"task_list_id"`
	Total       int    `json:"total"`
	Completed   int    `json:"completed"`
	InProgress  int    `json:"in_progress"`
	Pending     int    `json:"pending"`
	Cancelled   int    `json:"cancelled"`
	CompleteRate float64 `json:"complete_rate"`
}

// GetTaskStatsByTaskList 按清单统计任务完成情况。
func (s *Service) GetTaskStatsByTaskList(taskListID string) (*TaskListStats, error) {
	if _, err := s.store.GetTaskList(taskListID); err != nil {
		return nil, err
	}
	all := s.store.ListTasks()
	stats := &TaskListStats{TaskListID: taskListID}
	for _, t := range all {
		if t.TaskListID != taskListID {
			continue
		}
		stats.Total++
		switch t.Status {
		case model.TaskStatusDone:
			stats.Completed++
		case model.TaskStatusDoing:
			stats.InProgress++
		case model.TaskStatusTodo:
			stats.Pending++
		case model.TaskStatusCancelled:
			stats.Cancelled++
		}
	}
	if stats.Total > 0 {
		stats.CompleteRate = float64(stats.Completed) / float64(stats.Total)
	}
	return stats, nil
}
