package preferences

import "strings"

type Checker interface {
	Allow(string) bool
}

type requiredChecker struct{}

func (*requiredChecker) Allow(value string) bool { return strings.TrimSpace(value) != "" }

func NewChecker(enabled bool) Checker {
	if !enabled {
		return nil
	}
	return &requiredChecker{}
}
