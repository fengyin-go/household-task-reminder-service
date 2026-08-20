package preferences

import "strings"

type Checker interface {
	Allow(string) bool
}

type requiredChecker struct{}

func (c *requiredChecker) Allow(value string) bool {
	if c == nil {
		return true
	}
	return strings.TrimSpace(value) != ""
}

func NewChecker(enabled bool) Checker {
	if !enabled {
		var checker *requiredChecker
		return checker
	}
	return &requiredChecker{}
}
