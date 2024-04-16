package mutexes_concept

import (
	"fmt"
	"sync"
)

func increment_with_mutex(x *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	*x += 1
	mu.Unlock()
}

func decrement_with_mutex(x *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	*x -= 1
	mu.Unlock()
}

func Demo_With_Mutex() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	x := 0

	for i := 0; i < 200; i++ {
		wg.Add(2)
		go decrement_with_mutex(&x, &wg, &mu)
		go increment_with_mutex(&x, &wg, &mu)
	}

	wg.Wait()
	fmt.Println("x = ", x)
}
