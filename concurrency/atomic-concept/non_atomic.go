package atomic_concept

import (
	"fmt"
	"sync"
)

func NonAtomicDemo() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 100; c++ {
				ops++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)
}
