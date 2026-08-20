package service

import (
	"sort"
	"time"

	"todolist/internal/model"
	"todolist/pkg/idgen"
)

func (s *Service) CreateReminder(r model.Reminder) (*model.Reminder, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetTask(r.TaskID); err != nil {
		return nil, model.NewValidationError("task_id", "关联任务不存在")
	}
	r.ID = idgen.Hex()
	r.CreatedAt = time.Now()
	if err := s.store.CreateReminder(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *Service) GetReminder(id string) (*model.Reminder, error) {
	return s.store.GetReminder(id)
}

func (s *Service) ListReminders(filter model.ReminderFilter, page, size int) ([]*model.Reminder, int, error) {
	all := s.store.ListReminders()
	matched := make([]*model.Reminder, 0, len(all))
	for _, r := range all {
		if filter.Match(r) {
			matched = append(matched, r)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].RemindAt.Before(matched[j].RemindAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Reminder{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateReminder(id string, r model.Reminder) (*model.Reminder, error) {
	existing, err := s.store.GetReminder(id)
	if err != nil {
		return nil, err
	}
	if r.TaskID != "" {
		if _, err := s.store.GetTask(r.TaskID); err != nil {
			return nil, model.NewValidationError("task_id", "关联任务不存在")
		}
		existing.TaskID = r.TaskID
	}
	if !r.RemindAt.IsZero() {
		existing.RemindAt = r.RemindAt
	}
	if r.Message != "" {
		existing.Message = r.Message
	}
	if err := existing.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.UpdateReminder(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *Service) DeleteReminder(id string) error {
	return s.store.DeleteReminder(id)
}
