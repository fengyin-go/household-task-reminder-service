package service_test

import (
	"testing"
	"time"

	"todolist/internal/batch"
	"todolist/internal/service"
)

func TestFailedBatchItemDoesNotHangOrDropSuccessfulReminder(t *testing.T) {
	done := make(chan []batch.Result, 1)
	go func() {
		results, _ := (service.BatchReminderService{}).Dispatch([]batch.Item{{ID: "bad", Fail: true}, {ID: "good"}})
		done <- results
	}()
	select {
	case results := <-done:
		if len(results) != 1 || results[0].ID != "good" {
			t.Fatalf("successful reminder result lost: %#v", results)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("batch dispatch hung after a failed item")
	}
}
