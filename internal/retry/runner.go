package retry

import "errors"

type Runner struct {
	repo *Repository
}

func NewRunner(repo *Repository) *Runner { return &Runner{repo: repo} }

func (r *Runner) Deliver(id string) error {
	err := r.repo.Commit(id)
	if err == nil {
		return nil
	}
	var committed *CommitError
	if errors.As(err, &committed) && committed.Committed {
		return err
	}
	return r.repo.Commit(id)
}
