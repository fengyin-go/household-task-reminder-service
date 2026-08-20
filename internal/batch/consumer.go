package batch

func Consume(in <-chan Item, out chan<- Result) {
	for item := range in {
		out <- Result{ID: item.ID}
	}
}
