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
	s.tags[id] = tags
}

func (s *Store) Get(id string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.tags[id]
}
