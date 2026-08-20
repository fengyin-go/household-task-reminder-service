package model

import (
	"strings"
	"time"
)

// Subtask 任务下的子任务，用于拆解复杂任务。
type Subtask struct {
	ID        string    `json:"id"`
	TaskID    string    `json:"task_id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Subtask) Validate() error {
	s.Title = strings.TrimSpace(s.Title)
	if s.TaskID == "" {
		return NewValidationError("task_id", "所属任务不能为空")
	}
	if s.Title == "" {
		return NewValidationError("title", "子任务标题不能为空")
	}
	return nil
}

// SubtaskFilter 子任务筛选条件。
type SubtaskFilter struct {
	TaskID string
	Done   *bool
}

func (f SubtaskFilter) Match(s *Subtask) bool {
	if f.TaskID != "" && s.TaskID != f.TaskID {
		return false
	}
	if f.Done != nil && s.Done != *f.Done {
		return false
	}
	return true
}
