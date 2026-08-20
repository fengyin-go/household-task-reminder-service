package service_test

import (
	"testing"

	"todolist/internal/payload"
	"todolist/internal/service"
)

func TestQueuedReminderPayloadOwnsItsBytes(t *testing.T) {
	cache := payload.NewCache()
	svc := service.NewPayloadService(payload.NewProducer(payload.NewPool(), cache), payload.NewConsumer(cache))
	first := svc.Queue("one", "first reminder")
	svc.Queue("two", "second reminder")
	if string(first) != "first reminder" || svc.Read("one") != "first reminder" {
		t.Fatalf("first payload was overwritten: returned=%q cached=%q", first, svc.Read("one"))
	}
}
