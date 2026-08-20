package batch

func Consume(in <-chan Item, out chan<- Result) {
	for item := range in {
		if Empty(item) {
			continue
		}
		out <- Result{ID: item.ID}
	}
}
