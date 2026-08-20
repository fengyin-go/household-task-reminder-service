package summary

type Worker struct{ cache *Cache }

func NewWorker(cache *Cache) *Worker { return &Worker{cache: cache} }

func (w *Worker) Aggregate(id string, tags []string) {
	if id == "" {
		return
	}
	w.cache.Store(id, append([]string(nil), tags...))
}
