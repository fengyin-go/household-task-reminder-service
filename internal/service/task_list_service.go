package service

import (
	"sort"
	"time"

	"todolist/internal/model"
	"todolist/pkg/idgen"
)

func (s *Service) CreateTaskList(tl model.TaskList) (*model.TaskList, error) {
	if err := tl.Validate(); err != nil {
		return nil, err
	}
	tl.ID = idgen.Hex()
	tl.CreatedAt = time.Now()
	tl.UpdatedAt = tl.CreatedAt
	if err := s.store.CreateTaskList(&tl); err != nil {
		return nil, err
	}
	return &tl, nil
}

func (s *Service) GetTaskList(id string) (*model.TaskList, error) {
	return s.store.GetTaskList(id)
}

func (s *Service) ListTaskLists(filter model.TaskListFilter, page, size int) ([]*model.TaskList, int, error) {
	all := s.store.ListTaskLists()
	matched := make([]*model.TaskList, 0, len(all))
	for _, tl := range all {
		if filter.Match(tl) {
			matched = append(matched, tl)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.TaskList{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateTaskList(id string, tl model.TaskList) (*model.TaskList, error) {
	existing, err := s.store.GetTaskList(id)
	if err != nil {
		return nil, err
	}
	if tl.Name != "" {
		existing.Name = tl.Name
	}
	if tl.Description != "" {
		existing.Description = tl.Description
	}
	if err := existing.Validate(); err != nil {
		return nil, err
	}
	existing.UpdatedAt = time.Now()
	if err := s.store.UpdateTaskList(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *Service) DeleteTaskList(id string) error {
	return s.store.DeleteTaskList(id)
}
