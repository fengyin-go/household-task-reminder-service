package summary

type Worker struct{ cache *Cache }

func NewWorker(cache *Cache) *Worker { return &Worker{cache: cache} }

func (w *Worker) Aggregate(id string, tags []string) { w.cache.Store(id, tags) }
