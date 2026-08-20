package service_test

import (
	"testing"

	"todolist/internal/cleanup"
	"todolist/internal/service"
)

func TestPartialCleanupIsNotRetried(t *testing.T) {
	repo := &cleanup.Repository{Fail: true}
	svc := service.NewTaskListCleanup(cleanup.NewRetry(repo), nil)
	status, err := svc.Delete("list-1")
	if err == nil || status.Retryable {
		t.Fatalf("partial cleanup was hidden or marked retryable: err=%v retryable=%v", err, status.Retryable)
	}
	if repo.Cleanups != 1 {
		t.Fatalf("cleanup ran %d times after partial failure", repo.Cleanups)
	}
}
