package retry

import "errors"

type Repository struct {
	entries []string
	fail    bool
}

func NewRepository(failAfterCommit bool) *Repository {
	return &Repository{fail: failAfterCommit}
}

func (r *Repository) Commit(id string) error {
	r.entries = append(r.entries, id)
	if r.fail {
		r.fail = false
		return errors.New("delivery failed")
	}
	return nil
}

func (r *Repository) Count() int { return len(r.entries) }
