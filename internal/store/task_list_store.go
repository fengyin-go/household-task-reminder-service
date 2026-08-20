package store

import (
	"todolist/internal/model"
)

func (s *MemoryStore) CreateTaskList(tl *model.TaskList) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.taskLists[tl.ID]; ok {
		return ErrConflict
	}
	for _, exist := range s.taskLists {
		if exist.Name == tl.Name {
			return ErrConflict
		}
	}
	s.taskLists[tl.ID] = tl
	return nil
}

func (s *MemoryStore) GetTaskList(id string) (*model.TaskList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	tl, ok := s.taskLists[id]
	if !ok {
		return nil, ErrNotFound
	}
	return tl, nil
}

func (s *MemoryStore) ListTaskLists() []*model.TaskList {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.TaskList, 0, len(s.taskLists))
	for _, tl := range s.taskLists {
		list = append(list, tl)
	}
	return list
}

func (s *MemoryStore) UpdateTaskList(tl *model.TaskList) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.taskLists[tl.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.taskLists {
		if exist.ID != tl.ID && exist.Name == tl.Name {
			return ErrConflict
		}
	}
	s.taskLists[tl.ID] = tl
	return nil
}

func (s *MemoryStore) DeleteTaskList(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.taskLists[id]; !ok {
		return ErrNotFound
	}
	delete(s.taskLists, id)
	return nil
}
