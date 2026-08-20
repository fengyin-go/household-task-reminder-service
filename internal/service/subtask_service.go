package service

import (
	"sort"
	"time"

	"todolist/internal/model"
	"todolist/pkg/idgen"
)

// CreateSubtask 创建子任务。
func (s *Service) CreateSubtask(st model.Subtask) (*model.Subtask, error) {
	if err := st.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetTask(st.TaskID); err != nil {
		return nil, model.NewValidationError("task_id", "所属任务不存在")
	}
	st.ID = idgen.Hex()
	st.CreatedAt = time.Now()
	st.UpdatedAt = st.CreatedAt
	if err := s.store.CreateSubtask(&st); err != nil {
		return nil, err
	}
	return &st, nil
}

// ListSubtasks 分页列出子任务，支持按任务/完成状态筛选。
func (s *Service) ListSubtasks(filter model.SubtaskFilter, page, size int) ([]*model.Subtask, int, error) {
	all := s.store.ListSubtasks()
	matched := make([]*model.Subtask, 0, len(all))
	for _, st := range all {
		if filter.Match(st) {
			matched = append(matched, st)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Subtask{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

// ToggleSubtask 切换子任务完成状态。
func (s *Service) ToggleSubtask(id string) (*model.Subtask, error) {
	st, err := s.store.GetSubtask(id)
	if err != nil {
		return nil, err
	}
	st.Done = !st.Done
	st.UpdatedAt = time.Now()
	if err := s.store.UpdateSubtask(st); err != nil {
		return nil, err
	}
	return st, nil
}

// SetSubtaskDone 显式设置子任务完成状态。
func (s *Service) SetSubtaskDone(id string, done bool) (*model.Subtask, error) {
	st, err := s.store.GetSubtask(id)
	if err != nil {
		return nil, err
	}
	st.Done = done
	st.UpdatedAt = time.Now()
	if err := s.store.UpdateSubtask(st); err != nil {
		return nil, err
	}
	return st, nil
}

// DeleteSubtask 删除子任务。
func (s *Service) DeleteSubtask(id string) error {
	return s.store.DeleteSubtask(id)
}

// SubtaskStats 某任务子任务完成进度。
type SubtaskStats struct {
	TaskID   string  `json:"task_id"`
	Total    int     `json:"total"`
	Done     int     `json:"done"`
	Progress float64 `json:"progress"`
}

// GetSubtaskStats 返回任务子任务完成进度。
func (s *Service) GetSubtaskStats(taskID string) (*SubtaskStats, error) {
	if _, err := s.store.GetTask(taskID); err != nil {
		return nil, err
	}
	all := s.store.ListSubtasks()
	stats := &SubtaskStats{TaskID: taskID}
	for _, st := range all {
		if st.TaskID == taskID {
			stats.Total++
			if st.Done {
				stats.Done++
			}
		}
	}
	if stats.Total > 0 {
		stats.Progress = float64(stats.Done) / float64(stats.Total)
	}
	return stats, nil
}
