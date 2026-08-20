package store

import (
	"sync"

	"todolist/internal/model"
)

type MemoryStore struct {
	mu        sync.RWMutex
	tasks     map[string]*model.Task
	taskLists map[string]*model.TaskList
	tags      map[string]*model.Tag
	reminders map[string]*model.Reminder
	subtasks  map[string]*model.Subtask
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks:     make(map[string]*model.Task),
		taskLists: make(map[string]*model.TaskList),
		tags:      make(map[string]*model.Tag),
		reminders: make(map[string]*model.Reminder),
		subtasks:  make(map[string]*model.Subtask),
	}
}

var _ Store = (*MemoryStore)(nil)
