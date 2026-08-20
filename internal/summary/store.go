package summary

import "sync"

type Store struct {
	mu   sync.RWMutex
	tags map[string][]string
}

func NewStore() *Store { return &Store{tags: map[string][]string{}} }

func (s *Store) Set(id string, tags []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tags[id] = append([]string(nil), tags...)
}

func (s *Store) Has(id string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.tags[id]
	return ok
}

func (s *Store) Get(id string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]string(nil), s.tags[id]...)
}

func clone(tags []string) []string { return append([]string(nil), tags...) }
