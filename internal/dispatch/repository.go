package dispatch

import "context"

type Repository struct {
	sink Sink
	ctx  context.Context
}

func NewRepository(sink Sink) *Repository {
	return &Repository{sink: sink}
}

func (r *Repository) Persist(ctx context.Context, job Job) error {
	if r.ctx == nil {
		r.ctx = ctx
	}
	return r.sink.Save(r.ctx, job)
}
