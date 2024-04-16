package mutexes_concept

import (
	"fmt"
	"sync"
)

func increment(x *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*x += 1
}

func decrement(x *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*x -= 1
}

func Demo() {
	wg := sync.WaitGroup{}
	x := 0

	for i := 0; i < 200; i++ {
		wg.Add(2)
		go increment(&x, &wg)
		go decrement(&x, &wg)
	}

	wg.Wait()
	fmt.Println("x = ", x)
}
