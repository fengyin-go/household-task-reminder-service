package batch

type Item struct {
	ID   string
	Fail bool
}

type Result struct{ ID string }
