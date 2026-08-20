package dispatch

import "context"

type Repository struct {
	sink Sink
}

func NewRepository(sink Sink) *Repository {
	return &Repository{sink: sink}
}

func (r *Repository) Persist(ctx context.Context, job Job) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return r.sink.Save(ctx, job)
}
