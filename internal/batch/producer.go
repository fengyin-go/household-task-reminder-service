package batch

func Produce(item Item, out chan<- Item) {
	defer close(out)
	if item.Fail || !item.Valid() {
		return
	}
	out <- item
}

func Empty(item Item) bool {
	return item.Fail || !item.Valid()
}
