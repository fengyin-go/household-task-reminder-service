package events

type Worker struct{ repo *Repository }

func NewWorker(repo *Repository) *Worker { return &Worker{repo: repo} }

func (w *Worker) Apply(event Event) {
	if !event.Valid() {
		return
	}
	w.repo.Apply(event)
}
