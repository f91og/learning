package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 使用sync/atomic包的atomic counter来管理协程之间的状态
func TryAtomicCounters() {

	// use an unsigned integer to represent our (always-positive) counter.
	// 这里ops是我们要保证的在并发条件下是安全的计数器
	var ops uint64
	// A WaitGroup will help us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// start 50 goroutines that each increment the counter exactly 1000 times.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// 不能直接ops++, 因为递增操作不是个原子操作。利用atomic提供的AddUint64()原子操作，这个包里还提供了一些其他的原子操作
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// It’s safe to access ops now because we know no other goroutine is writing to it. Reading atomics safely while they are being updated is also possible, using functions like atomic.LoadUint64.
	// 因为ops在并发条件下是安全的，所以这里不管运行多少次结果都是50000
	fmt.Println("ops:", ops)

}
