package retry

import "errors"

var ErrTemporary = errors.New("temporary gateway failure")

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
		return &CommitError{Committed: true, Cause: ErrTemporary}
	}
	return nil
}

func (r *Repository) Count() int { return len(r.entries) }
