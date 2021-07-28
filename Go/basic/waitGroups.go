package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple goroutines to finish, we can use a wait group.

// a WaitGroup must be passed to functions by pointer
func worker3(id int, wg *sync.WaitGroup) {
	// defer是用来延迟执行函数的，而且延迟发生在函数 return 之后, 即下面的代码是让worker执行完毕了之后再执行wg.Done()
	// 利用defer可以方便安全的关闭资源，不必担心中间异常或者有地方没写资源释放语句而导致的资源未及时释放
	defer wg.Done() // Done()方法将WaitGroup的counter计数器减1

	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func TryWaitGroups() {
	// 这里时如何实例化一个WaitGroup的？貌似直接声明一个就开始用了？👈 因为WaitGroup是一个结构体，结构体不需要像channel那样先make分配内存再使用，
	// 而且go中也没有实例化这个概念，其实go在语法层面上更像C
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker3(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
	wg.Wait()
}
