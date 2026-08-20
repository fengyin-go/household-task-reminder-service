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

func (e Event) Advances(current TaskView) bool { return true }

func (e Event) Valid() bool { return true }
