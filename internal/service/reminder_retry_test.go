package service_test

import (
	"testing"

	"todolist/internal/retry"
	"todolist/internal/service"
)

func TestCommittedFailureDoesNotRetryReminder(t *testing.T) {
	repo := retry.NewRepository(true)
	delivery := service.NewReminderDelivery(retry.NewRunner(repo))
	result, err := delivery.Send("reminder-1")
	if err == nil {
		t.Fatal("partial commit must remain visible to the caller")
	}
	if result.Retryable {
		t.Fatal("already committed reminder was incorrectly marked retryable")
	}
	if repo.Count() != 1 {
		t.Fatalf("committed reminder count = %d, want exactly one side effect", repo.Count())
	}
}
