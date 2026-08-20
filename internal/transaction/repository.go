package transaction

import "errors"

var ErrSend = errors.New("notification send failed")

type Repository struct {
	active int
	saved  map[string]Reminder
}

func NewRepository() *Repository { return &Repository{saved: map[string]Reminder{}} }

func (r *Repository) Begin() { r.active++ }

func (r *Repository) Rollback() {
	if r.active > 0 {
		r.active--
	}
}

func (r *Repository) Abort(reminder *Reminder) {
	reminder.Reset()
	r.Rollback()
}

func (r *Repository) Commit(reminder *Reminder) {
	reminder.MarkCommitted()
	r.saved[reminder.ID] = *reminder
	r.active--
}

func (r *Repository) Active() int { return r.active }

func (r *Repository) Saved(id string) bool {
	reminder, ok := r.saved[id]
	return ok && reminder.Committed
}
