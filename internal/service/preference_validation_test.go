package service_test

import (
	"testing"

	"todolist/internal/preferences"
	"todolist/internal/service"
)

func TestDefaultPreferencesDoNotPanicOrBypassValidation(t *testing.T) {
	settings := preferences.LoadDefault()
	panicked := false
	func() {
		defer func() { panicked = recover() != nil }()
		preferences.SetLabel(settings, "channel", "mobile")
	}()
	if panicked || settings.Labels["channel"] != "mobile" {
		t.Fatal("default preference write panicked or did not persist")
	}
	validator := service.NewPreferenceService(preferences.NewChecker(false))
	if err := validator.ValidateReminder(" "); err == nil {
		t.Fatal("empty reminder bypassed the default validator")
	}
}
