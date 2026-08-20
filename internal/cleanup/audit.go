package cleanup

type Audit struct{ partial int }

func (a *Audit) Record(err error) {
	if err != nil {
		a.partial++
	}
}

func (a *Audit) PartialFailures() int { return a.partial }
