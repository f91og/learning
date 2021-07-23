package main

import "fmt"

func TryCloseChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// the more value will be false if jobs has been closed and all values in the channel have already been received
			// 看来channel除了返回基本的存储在其中的数据之外，还返回其他一些状态标志位
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j < 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // 利用内置的函数close来关闭channel
	fmt.Println("sent all jobs")

	<-done // 阻塞主协程以等待子协程的执行完毕。We await the worker using the synchronization approach we saw earlier.
}
