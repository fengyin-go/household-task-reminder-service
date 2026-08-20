package request

import (
	"context"
	"sync"
)

type Store struct {
	mu    sync.Mutex
	saved []string
	ctx   context.Context
}

func (s *Store) Save(ctx context.Context, id string) error {
	if s.ctx == nil {
		s.ctx = ctx
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
