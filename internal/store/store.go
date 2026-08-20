// Package store 定义数据访问接口与内存实现。
package store

import (
	"errors"

	"todolist/internal/model"
)

var (
	ErrNotFound = errors.New("记录不存在")
	ErrConflict = errors.New("记录已存在或状态冲突")
)

// Store 聚合全部实体的数据访问方法，便于测试时替换实现。
type Store interface {
	// Task
	CreateTask(t *model.Task) error
	GetTask(id string) (*model.Task, error)
	ListTasks() []*model.Task
	UpdateTask(t *model.Task) error
	DeleteTask(id string) error

	// TaskList
	CreateTaskList(tl *model.TaskList) error
	GetTaskList(id string) (*model.TaskList, error)
	ListTaskLists() []*model.TaskList
	UpdateTaskList(tl *model.TaskList) error
	DeleteTaskList(id string) error

	// Tag
	CreateTag(tag *model.Tag) error
	GetTag(id string) (*model.Tag, error)
	GetTagByName(name string) (*model.Tag, error)
	ListTags() []*model.Tag
	UpdateTag(tag *model.Tag) error
	DeleteTag(id string) error

	// Reminder
	CreateReminder(r *model.Reminder) error
	GetReminder(id string) (*model.Reminder, error)
	ListReminders() []*model.Reminder
	UpdateReminder(r *model.Reminder) error
	DeleteReminder(id string) error

	// Subtask
	CreateSubtask(s *model.Subtask) error
	GetSubtask(id string) (*model.Subtask, error)
	ListSubtasks() []*model.Subtask
	UpdateSubtask(s *model.Subtask) error
	DeleteSubtask(id string) error
}
