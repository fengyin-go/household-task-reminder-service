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
	s.mu.Lock()
	defer s.mu.Unlock()
	// 按请求自身的 ctx 判断取消，不 latch 到结构体字段，避免上一条请求的取消状态串到下一条。
	if err := ctx.Err(); err != nil {
		return err
	}
	s.saved = append(s.saved, id)
	return nil
}

func (s *Store) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.saved)
}
