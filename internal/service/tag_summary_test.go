package service_test

import (
	"testing"

	"todolist/internal/service"
	"todolist/internal/summary"
)

func TestSummaryUsesStableTagSnapshot(t *testing.T) {
	store := summary.NewStore()
	cache := summary.NewCache()
	svc := service.NewTagSummaryService(store, summary.NewWorker(cache))
	backing := make([]string, 1, 4)
	backing[0] = "home"
	store.Set("today", backing)
	svc.Refresh("today")
	backing = append(backing[:0], "work", "urgent")
	store.Set("today", backing)
	if cache.First("today") != "home" {
		t.Fatalf("summary observed later tag mutation: first=%q", cache.First("today"))
	}
}
