package request

import "context"

type Worker struct{ store *Store }

func NewWorker(store *Store) *Worker { return &Worker{store: store} }

func (w *Worker) Persist(ctx context.Context, id string) error {
	if !Active(ctx) {
		return context.Canceled
	}
	return w.store.Save(ctx, id)
}
