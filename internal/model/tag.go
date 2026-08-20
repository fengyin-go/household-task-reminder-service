package model

import (
	"strings"
	"time"
)

// Tag 标签实体。
type Tag struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Tag) Validate() error {
	t.Name = strings.TrimSpace(t.Name)
	if t.Name == "" {
		return NewValidationError("name", "标签名称不能为空")
	}
	return nil
}

// TagFilter 标签筛选条件。
type TagFilter struct {
	Keyword string
}

func (f TagFilter) Match(t *Tag) bool {
	if k := strings.ToLower(strings.TrimSpace(f.Keyword)); k != "" {
		if !strings.Contains(strings.ToLower(t.Name), k) {
			return false
		}
	}
	return true
}
