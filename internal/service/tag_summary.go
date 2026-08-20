package service

import "todolist/internal/summary"

type TagSummaryService struct {
	store  *summary.Store
	worker *summary.Worker
}

func NewTagSummaryService(store *summary.Store, worker *summary.Worker) *TagSummaryService {
	return &TagSummaryService{store: store, worker: worker}
}

func (s *TagSummaryService) Refresh(id string) {
	if !s.store.Has(id) {
		return
	}
	snapshot := s.store.Get(id)
	s.worker.Aggregate(id, append([]string(nil), snapshot...))
}
