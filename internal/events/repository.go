package events

type Repository struct {
	views map[string]TaskView
}

func NewRepository() *Repository { return &Repository{views: map[string]TaskView{}} }

func (r *Repository) Load(id string) TaskView { return r.views[id] }

func (r *Repository) Save(id string, view TaskView) { r.views[id] = view }

func (r *Repository) Apply(event Event) bool {
	current := r.Load(event.TaskID)
	if !event.Advances(current) {
		return false
	}
	r.Save(event.TaskID, TaskView{State: event.State, Version: event.Version})
	return true
}

func (r *Repository) Completed() int {
	n := 0
	for _, view := range r.views {
		if view.State != StateDone {
			continue
		}
		n++
	}
	return n
}
