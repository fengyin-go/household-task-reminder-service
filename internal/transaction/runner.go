package transaction

type Runner struct{ repo *Repository }

func NewRunner(repo *Repository) *Runner { return &Runner{repo: repo} }

func (r *Runner) Deliver(reminder *Reminder, fail bool) error {
	r.repo.Begin()
	if fail {
		r.repo.Abort(reminder)
		return ErrSend
	}
	r.repo.Commit(reminder)
	return nil
}

func (r *Runner) Active() int { return r.repo.Active() }
