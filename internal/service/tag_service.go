package service

import (
	"sort"
	"time"

	"todolist/internal/model"
	"todolist/pkg/idgen"
)

func (s *Service) CreateTag(tag model.Tag) (*model.Tag, error) {
	if err := tag.Validate(); err != nil {
		return nil, err
	}
	tag.ID = idgen.Hex()
	tag.CreatedAt = time.Now()
	if err := s.store.CreateTag(&tag); err != nil {
		return nil, err
	}
	return &tag, nil
}

func (s *Service) GetTag(id string) (*model.Tag, error) {
	return s.store.GetTag(id)
}

func (s *Service) ListTags(filter model.TagFilter, page, size int) ([]*model.Tag, int, error) {
	all := s.store.ListTags()
	matched := make([]*model.Tag, 0, len(all))
	for _, t := range all {
		if filter.Match(t) {
			matched = append(matched, t)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Tag{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateTag(id string, tag model.Tag) (*model.Tag, error) {
	existing, err := s.store.GetTag(id)
	if err != nil {
		return nil, err
	}
	if tag.Name != "" {
		existing.Name = tag.Name
	}
	if tag.Color != "" {
		existing.Color = tag.Color
	}
	if err := existing.Validate(); err != nil {
		return nil, err
	}
	if err := s.store.UpdateTag(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *Service) DeleteTag(id string) error {
	return s.store.DeleteTag(id)
}
