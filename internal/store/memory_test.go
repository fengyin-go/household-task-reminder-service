package store

import (
	"testing"
	"time"

	"todolist/internal/model"
)

func TestTaskCRUD(t *testing.T) {
	s := NewMemoryStore()
	task := &model.Task{ID: "t1", Title: "测试任务", TaskListID: "l1", Status: model.TaskStatusTodo}
	if err := s.CreateTask(task); err != nil {
		t.Fatalf("create: %v", err)
	}
	if err := s.CreateTask(&model.Task{ID: "t1", Title: "重复"}); err != ErrConflict {
		t.Fatalf("want conflict, got %v", err)
	}
	got, err := s.GetTask("t1")
	if err != nil || got.Title != "测试任务" {
		t.Fatalf("get: %v %v", got, err)
	}
	task.Status = model.TaskStatusDoing
	if err := s.UpdateTask(task); err != nil {
		t.Fatalf("update: %v", err)
	}
	if err := s.DeleteTask("t1"); err != nil {
		t.Fatalf("delete: %v", err)
	}
	if _, err := s.GetTask("t1"); err != ErrNotFound {
		t.Fatalf("want not found, got %v", err)
	}
}

func TestTaskListUniqueName(t *testing.T) {
	s := NewMemoryStore()
	tl := &model.TaskList{ID: "l1", Name: "清单A"}
	if err := s.CreateTaskList(tl); err != nil {
		t.Fatalf("create: %v", err)
	}
	if err := s.CreateTaskList(&model.TaskList{ID: "l2", Name: "清单A"}); err != ErrConflict {
		t.Fatalf("want conflict, got %v", err)
	}
	tl.Name = "清单B"
	if err := s.UpdateTaskList(tl); err != nil {
		t.Fatalf("update: %v", err)
	}
	if err := s.DeleteTaskList("l1"); err != nil {
		t.Fatalf("delete: %v", err)
	}
}

func TestTagUniqueName(t *testing.T) {
	s := NewMemoryStore()
	tag := &model.Tag{ID: "g1", Name: "工作"}
	if err := s.CreateTag(tag); err != nil {
		t.Fatalf("create: %v", err)
	}
	if err := s.CreateTag(&model.Tag{ID: "g2", Name: "工作"}); err != ErrConflict {
		t.Fatalf("want conflict, got %v", err)
	}
	byName, err := s.GetTagByName("工作")
	if err != nil || byName.ID != "g1" {
		t.Fatalf("get by name: %v %v", byName, err)
	}
	if err := s.DeleteTag("g1"); err != nil {
		t.Fatalf("delete: %v", err)
	}
}

func TestReminderCRUD(t *testing.T) {
	s := NewMemoryStore()
	r := &model.Reminder{ID: "r1", TaskID: "t1", RemindAt: time.Now(), Message: "提醒"}
	if err := s.CreateReminder(r); err != nil {
		t.Fatalf("create: %v", err)
	}
	if err := s.CreateReminder(&model.Reminder{ID: "r1"}); err != ErrConflict {
		t.Fatalf("want conflict, got %v", err)
	}
	got, err := s.GetReminder("r1")
	if err != nil || got.Message != "提醒" {
		t.Fatalf("get: %v %v", got, err)
	}
	if err := s.DeleteReminder("r1"); err != nil {
		t.Fatalf("delete: %v", err)
	}
}

func TestUpdateNotFound(t *testing.T) {
	s := NewMemoryStore()
	if err := s.UpdateTask(&model.Task{ID: "nope", Title: "x"}); err != ErrNotFound {
		t.Fatalf("task: want not found, got %v", err)
	}
	if err := s.UpdateTaskList(&model.TaskList{ID: "nope", Name: "x"}); err != ErrNotFound {
		t.Fatalf("tasklist: want not found, got %v", err)
	}
	if err := s.UpdateTag(&model.Tag{ID: "nope", Name: "x"}); err != ErrNotFound {
		t.Fatalf("tag: want not found, got %v", err)
	}
	if err := s.UpdateReminder(&model.Reminder{ID: "nope"}); err != ErrNotFound {
		t.Fatalf("reminder: want not found, got %v", err)
	}
}

func TestDeleteNotFound(t *testing.T) {
	s := NewMemoryStore()
	if err := s.DeleteTask("nope"); err != ErrNotFound {
		t.Fatalf("task: %v", err)
	}
	if err := s.DeleteTaskList("nope"); err != ErrNotFound {
		t.Fatalf("tasklist: %v", err)
	}
	if err := s.DeleteTag("nope"); err != ErrNotFound {
		t.Fatalf("tag: %v", err)
	}
	if err := s.DeleteReminder("nope"); err != ErrNotFound {
		t.Fatalf("reminder: %v", err)
	}
}
