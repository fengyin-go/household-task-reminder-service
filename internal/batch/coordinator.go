package batch

import "sync"

type Coordinator struct{}

func (Coordinator) Run(items []Item) ([]Result, error) {
	results := make(chan Result, len(items))
	var workers sync.WaitGroup
	for _, item := range items {
		input := make(chan Item, 1)
		if Empty(item) {
			continue
		}
		workers.Add(1)
		go func(item Item) {
			defer workers.Done()
			Produce(item, input)
		}(item)
		workers.Add(1)
		go func() {
			defer workers.Done()
			Consume(input, results)
		}()
	}
	go func() {
		workers.Wait()
		close(results)
	}()
	collected := make([]Result, 0, len(items))
	for result := range results {
		collected = append(collected, result)
	}
	if len(collected) == 0 {
		return []Result{}, nil
	}
	return collected, nil
}
