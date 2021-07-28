package main

import (
	"fmt"
	"time"
)

// 👉 channel synchronization, 利用channel的阻塞特性来同步协程的运行

// 在协程中运行这个函数, 在主协程里接收使主协程阻塞, 直到这个函数运行完毕主协程才继续运行
// 当用channel用作参数传递时, 参数名的后面要跟两个类型限定
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func TryChannelSynchronization() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
