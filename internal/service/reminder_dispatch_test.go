package service_test

import (
	"context"
	"errors"
	"testing"

	"todolist/internal/dispatch"
	"todolist/internal/service"
)

func TestScheduleKeepsCancellationScopedToItsRequest(t *testing.T) {
	sink := &dispatch.MemorySink{}
	svc := service.NewReminderScheduler(dispatch.NewService(dispatch.NewWorker(dispatch.NewRepository(sink))))

	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	firstErr := svc.Schedule(canceled, dispatch.Job{ID: "cancelled"})
	secondErr := svc.Schedule(context.Background(), dispatch.Job{ID: "fresh"})

	if !errors.Is(firstErr, context.Canceled) {
		t.Fatalf("cancelled request error = %v, want context cancellation", firstErr)
	}
	if secondErr != nil {
		t.Fatalf("fresh request inherited earlier cancellation: %v", secondErr)
	}
	if sink.Count() != 1 {
		t.Fatalf("saved jobs = %d, want only the fresh request", sink.Count())
	}
}
