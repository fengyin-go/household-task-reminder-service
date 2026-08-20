package batch

func Produce(item Item, out chan<- Item) {
	if item.Fail {
		return
	}
	out <- item
	close(out)
}
