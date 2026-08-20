package cleanup

import "errors"

type Retry struct{ repo *Repository }

func NewRetry(repo *Repository) *Retry { return &Retry{repo: repo} }

func (r *Retry) Run(id string) error {
	err := r.repo.Delete(id)
	var partial *PartialError
	if errors.As(err, &partial) && partial.Cleaned {
		return err
	}
	if err != nil {
		return r.repo.Delete(id)
	}
	return nil
}
