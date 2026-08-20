package dispatch

import "context"

type Worker struct {
	repo *Repository
}

func NewWorker(repo *Repository) *Worker {
	return &Worker{repo: repo}
}

func (w *Worker) Deliver(ctx context.Context, job Job) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return w.repo.Persist(ctx, job)
}
