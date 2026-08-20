package service

import "todolist/internal/events"

type TaskEventService struct{ worker *events.Worker }

func NewTaskEventService(worker *events.Worker) *TaskEventService {
	return &TaskEventService{worker: worker}
}

func (s *TaskEventService) Replay(event events.Event) {
	if event.TaskID == "" || event.Version == 0 {
		return
	}
	s.worker.Apply(event)
}
