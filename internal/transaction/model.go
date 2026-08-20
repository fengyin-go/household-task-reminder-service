package transaction

type Reminder struct {
	ID        string
	Committed bool
}

func (r *Reminder) MarkCommitted() { r.Committed = true }
