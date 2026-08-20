package dispatch

import "context"

type Service struct {
	worker *Worker
}

func NewService(worker *Worker) *Service {
	return &Service{worker: worker}
}

func (s *Service) Schedule(ctx context.Context, job Job) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return s.worker.Deliver(ctx, job)
}
