package events

const (
	StateDoing = "doing"
	StateDone  = "done"
)

type Event struct {
	TaskID  string
	State   string
	Version int
}

type TaskView struct {
	State   string
	Version int
}

func (e Event) Advances(current TaskView) bool {
	return e.State != "" && e.Version > current.Version
}

func (e Event) Valid() bool {
	return e.TaskID != "" && e.Version > 0 && (e.State == StateDoing || e.State == StateDone)
}
