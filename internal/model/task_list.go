package model

import (
	"strings"
	"time"
)

// TaskList 清单实体。
type TaskList struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (tl *TaskList) Validate() error {
	tl.Name = strings.TrimSpace(tl.Name)
	if tl.Name == "" {
		return NewValidationError("name", "清单名称不能为空")
	}
	return nil
}

// TaskListFilter 清单筛选条件。
type TaskListFilter struct {
	Keyword string
}

func (f TaskListFilter) Match(tl *TaskList) bool {
	if k := strings.ToLower(strings.TrimSpace(f.Keyword)); k != "" {
		if !strings.Contains(strings.ToLower(tl.Name), k) &&
			!strings.Contains(strings.ToLower(tl.Description), k) {
			return false
		}
	}
	return true
}
