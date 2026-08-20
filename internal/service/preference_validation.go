package service

import (
	"errors"

	"todolist/internal/preferences"
)

var ErrInvalidReminder = errors.New("invalid reminder")

type PreferenceService struct{ checker preferences.Checker }

func NewPreferenceService(checker preferences.Checker) *PreferenceService {
	return &PreferenceService{checker: checker}
}

func (s *PreferenceService) ValidateReminder(message string) error {
	if s.checker == nil {
		return nil
	}
	if !s.checker.Allow(message) {
		return ErrInvalidReminder
	}
	return nil
}
