package retry

import "errors"

var ErrTemporary = errors.New("temporary gateway failure")
