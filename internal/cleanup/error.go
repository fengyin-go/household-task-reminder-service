package cleanup

import "errors"

var ErrDeleteFailed = errors.New("task list delete failed")

type PartialError struct {
	Cleaned bool
	Cause   error
}

func (e *PartialError) Error() string { return e.Cause.Error() }
func (e *PartialError) Unwrap() error { return e.Cause }
