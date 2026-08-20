package store

import (
	"strings"

	"todolist/internal/model"
)

func (s *MemoryStore) CreateTag(tag *model.Tag) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tags[tag.ID]; ok {
		return ErrConflict
	}
	for _, exist := range s.tags {
		if strings.EqualFold(exist.Name, tag.Name) {
			return ErrConflict
		}
	}
	s.tags[tag.ID] = tag
	return nil
}

func (s *MemoryStore) GetTag(id string) (*model.Tag, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	tag, ok := s.tags[id]
	if !ok {
		return nil, ErrNotFound
	}
	return tag, nil
}

func (s *MemoryStore) GetTagByName(name string) (*model.Tag, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, tag := range s.tags {
		if strings.EqualFold(tag.Name, name) {
			return tag, nil
		}
	}
	return nil, ErrNotFound
}

func (s *MemoryStore) ListTags() []*model.Tag {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Tag, 0, len(s.tags))
	for _, tag := range s.tags {
		list = append(list, tag)
	}
	return list
}

func (s *MemoryStore) UpdateTag(tag *model.Tag) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tags[tag.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.tags {
		if exist.ID != tag.ID && strings.EqualFold(exist.Name, tag.Name) {
			return ErrConflict
		}
	}
	s.tags[tag.ID] = tag
	return nil
}

func (s *MemoryStore) DeleteTag(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tags[id]; !ok {
		return ErrNotFound
	}
	delete(s.tags, id)
	return nil
}
