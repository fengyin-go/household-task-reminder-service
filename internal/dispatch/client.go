package dispatch

import (
	"context"
	"sync"
)

type Sink interface {
	Save(context.Context, Job) error
}

type MemorySink struct {
	mu    sync.Mutex
	Saved []string
}

func (s *MemorySink) Save(ctx context.Context, job Job) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Saved = append(s.Saved, job.ID)
	return nil
}

func (s *MemorySink) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.Saved)
}
