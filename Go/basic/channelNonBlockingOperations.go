package main

import "fmt"

// implement non-blocking sends, receives, and even non-blocking multi-way selects.

func TryNonBlockingChannelOperations() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received msg:", msg)
	default:
		fmt.Println("no msg received")
	}

	msg := "hi"
	select {
	case messages <- msg: // select会随机挑选一个case尝试取执行，如果这个case不阻塞则能成功执行case里的语句
		fmt.Println("sent msg:", msg)
	default:
		fmt.Println("no msg sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received msg:", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
