package cleanup

type Retry struct{ repo *Repository }

func NewRetry(repo *Repository) *Retry { return &Retry{repo: repo} }

func (r *Retry) Run(id string) error {
	if err := r.repo.Delete(id); err != nil {
		return r.repo.Delete(id)
	}
	return nil
}
