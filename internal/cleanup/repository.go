package cleanup

type Repository struct {
	Cleanups int
	Fail     bool
}

func (r *Repository) Delete(id string) error {
	r.Cleanups++
	if r.Fail {
		r.Fail = false
		return ErrDeleteFailed
	}
	return nil
}
