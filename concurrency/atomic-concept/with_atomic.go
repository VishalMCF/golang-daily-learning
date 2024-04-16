package atomic_concept

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
func AddInt64(address *int64, delta int64)(new int64)
func LoadInt64(address *int64) (value int64)
func StoreInt64(address *int64, value int64)
func SwapInt64(address *int64, new int64) (old int64)
func CompareAndSwapInt64(address *int64, old, new int64) (swapped bool)
*/

func WithAtomicDemo() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 100; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops.Load())
}

//func WithAtomicDemo() {
//	var ops uint64
//	var wg sync.WaitGroup
//
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go func() {
//			for c := 0; c < 100; c++ {
//				atomic.AddUint64(&ops, 1)
//			}
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//	fmt.Println("ops:", atomic.LoadUint64(&ops))
//}
