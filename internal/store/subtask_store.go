package store

import "todolist/internal/model"

// CreateSubtask 创建子任务。
func (s *MemoryStore) CreateSubtask(st *model.Subtask) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.subtasks[st.ID]; ok {
		return ErrConflict
	}
	s.subtasks[st.ID] = st
	return nil
}

// GetSubtask 按 ID 查询子任务。
func (s *MemoryStore) GetSubtask(id string) (*model.Subtask, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	st, ok := s.subtasks[id]
	if !ok {
		return nil, ErrNotFound
	}
	return st, nil
}

// ListSubtasks 列出全部子任务。
func (s *MemoryStore) ListSubtasks() []*model.Subtask {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Subtask, 0, len(s.subtasks))
	for _, st := range s.subtasks {
		list = append(list, st)
	}
	return list
}

// UpdateSubtask 更新子任务。
func (s *MemoryStore) UpdateSubtask(st *model.Subtask) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.subtasks[st.ID]; !ok {
		return ErrNotFound
	}
	s.subtasks[st.ID] = st
	return nil
}

// DeleteSubtask 删除子任务。
func (s *MemoryStore) DeleteSubtask(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.subtasks[id]; !ok {
		return ErrNotFound
	}
	delete(s.subtasks, id)
	return nil
}
