package main

import (
	"fmt"
	"time"
)

// 通过go的select来等待多个channel操作

func TrySelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// select会同时等待所有case，如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行，所以这里我们循环了两次，因为需要等待两个通道返回结果
	// 和switch一样可以用default语句。在没有case可以执行的时候，如果有defaul语句，那么select会执行default语句，否则select会一直阻塞，直到某个通信可以运行
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

}
