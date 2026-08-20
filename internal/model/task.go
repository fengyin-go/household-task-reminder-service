package model

import (
	"strings"
	"time"
)

const (
	TaskStatusTodo      = "todo"
	TaskStatusDoing     = "doing"
	TaskStatusDone      = "done"
	TaskStatusCancelled = "cancelled"
)

var taskTransitions = map[string]map[string]bool{
	TaskStatusTodo:      {TaskStatusDoing: true, TaskStatusCancelled: true},
	TaskStatusDoing:     {TaskStatusDone: true, TaskStatusCancelled: true},
	TaskStatusDone:      {TaskStatusTodo: true},
	TaskStatusCancelled: {TaskStatusTodo: true},
}

// CanTransitionTask 校验任务状态流转是否合法。
func CanTransitionTask(from, to string) bool {
	if m, ok := taskTransitions[from]; ok {
		return m[to]
	}
	return false
}

// Task 任务实体。
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	TaskListID  string    `json:"task_list_id"`
	Priority    int       `json:"priority"`
	DueDate     time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Task) Validate() error {
	t.Title = strings.TrimSpace(t.Title)
	if t.Title == "" {
		return NewValidationError("title", "任务标题不能为空")
	}
	if t.TaskListID == "" {
		return NewValidationError("task_list_id", "所属清单不能为空")
	}
	if t.Status == "" {
		t.Status = TaskStatusTodo
	}
	if t.Status != TaskStatusTodo && t.Status != TaskStatusDoing && t.Status != TaskStatusDone && t.Status != TaskStatusCancelled {
		return NewValidationError("status", "任务状态不合法")
	}
	if t.Priority < 0 {
		return NewValidationError("priority", "优先级不能为负数")
	}
	return nil
}

// TaskFilter 任务筛选条件。
type TaskFilter struct {
	TaskListID string
	Status     string
	Keyword    string
}

func (f TaskFilter) Match(t *Task) bool {
	if f.TaskListID != "" && t.TaskListID != f.TaskListID {
		return false
	}
	if f.Status != "" && t.Status != f.Status {
		return false
	}
	if k := strings.ToLower(strings.TrimSpace(f.Keyword)); k != "" {
		if !strings.Contains(strings.ToLower(t.Title), k) &&
			!strings.Contains(strings.ToLower(t.Description), k) {
			return false
		}
	}
	return true
}
