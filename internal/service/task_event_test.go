package service_test

import (
	"testing"

	"todolist/internal/events"
	"todolist/internal/service"
)

func TestOldReminderEventCannotReopenCompletedTask(t *testing.T) {
	repo := events.NewRepository()
	svc := service.NewTaskEventService(events.NewWorker(repo))
	svc.Replay(events.Event{TaskID: "task-1", State: events.StateDone, Version: 2})
	svc.Replay(events.Event{TaskID: "task-1", State: events.StateDoing, Version: 1})
	view := repo.Load("task-1")
	if view.State != events.StateDone || repo.Completed() != 1 {
		t.Fatalf("stale event reopened task: state=%s completed=%d", view.State, repo.Completed())
	}
}
