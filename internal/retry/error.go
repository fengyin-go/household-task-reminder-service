package retry

import "fmt"

// CommitError records whether a delivery reached its durable side effect.
type CommitError struct {
	Committed bool
	Cause     error
}

func (e *CommitError) Error() string {
	return fmt.Sprintf("delivery commit failed: %v", e.Cause)
}

func (e *CommitError) Unwrap() error { return e.Cause }
