package model

import (
	"strings"
	"time"
)

// Reminder 提醒实体。
type Reminder struct {
	ID        string    `json:"id"`
	TaskID    string    `json:"task_id"`
	RemindAt  time.Time `json:"remind_at"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *Reminder) Validate() error {
	if r.TaskID == "" {
		return NewValidationError("task_id", "关联任务不能为空")
	}
	if r.RemindAt.IsZero() {
		return NewValidationError("remind_at", "提醒时间不能为空")
	}
	r.Message = strings.TrimSpace(r.Message)
	return nil
}

// ReminderFilter 提醒筛选条件。
type ReminderFilter struct {
	TaskID string
}

func (f ReminderFilter) Match(r *Reminder) bool {
	if f.TaskID != "" && r.TaskID != f.TaskID {
		return false
	}
	return true
}
