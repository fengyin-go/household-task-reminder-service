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
	s.worker.Aggregate(id, s.store.Get(id))
}
