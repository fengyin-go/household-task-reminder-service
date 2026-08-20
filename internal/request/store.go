package request

import (
	"context"
	"sync"
)

type Store struct {
	mu    sync.Mutex
	saved []string
}

func (s *Store) Save(ctx context.Context, id string) error {
	if !Active(ctx) {
		return context.Canceled
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.saved = append(s.saved, id)
	return nil
}

func (s *Store) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.saved)
}
