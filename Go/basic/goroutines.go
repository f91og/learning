// A goroutine is a lightweight thread of execution.
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func TryGoroutines() {
	f("direct")

	// 执行函数的时候在前面加上go让其在新协程中执行
	go f("goroutine")

	// go中函数里面可以定义函数，这里不光定义了函数，还定义后立马就执行了
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
