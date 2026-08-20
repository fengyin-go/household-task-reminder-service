package service_test

import (
	"context"
	"errors"
	"testing"

	"todolist/internal/request"
	"todolist/internal/service"
)

func TestTimedOutReminderDoesNotPersistOrPoisonNextRequest(t *testing.T) {
	store := &request.Store{}
	svc := service.NewReminderRequestService(request.NewClient(request.NewWorker(store)))
	expired, cancel := context.WithCancel(context.Background())
	cancel()
	first := svc.Create(expired, "timed-out")
	second := svc.Create(context.Background(), "fresh")
	if !errors.Is(first, context.Canceled) || store.Count() != 1 {
		t.Fatalf("cancelled request persisted or lost cancellation: err=%v saved=%d", first, store.Count())
	}
	if second != nil {
		t.Fatalf("fresh request inherited a stale deadline: %v", second)
	}
}
