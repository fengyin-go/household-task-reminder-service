package retry

type Runner struct {
	repo *Repository
}

func NewRunner(repo *Repository) *Runner { return &Runner{repo: repo} }

func (r *Runner) Deliver(id string) error {
	if err := r.repo.Commit(id); err != nil {
		return r.repo.Commit(id)
	}
	return nil
}
