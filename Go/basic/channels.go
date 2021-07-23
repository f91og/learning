// Channels are the pipes that connect concurrent goroutines.
package main

import (
	"fmt"
)

func TryChannels() {
	// 👉 使用无缓冲区的channel
	// 通过make来实例化一个channel，需要指定channel里具体接受什么类型的值
	messages := make(chan string)

	go func() { messages <- "ping" }() // 这个协程里里面就算阻塞了, main函数(goroutine 1)还是能正常执行

	// By default sends and receives block until both the sender and receiver are ready. 默认情况下发送方和接收方都会阻塞
	// This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization
	// 测试1：将go func() { messages <- "ping" }() 注释掉后运行发现下面的执行报fatal error: all goroutines are asleep - deadlock! 下面的channel读取会一直等待
	// 测试2：不在另一个协程中向通道发信息，就在本地直接发 messages <- "ping"，会报和上面一样的错误。看来通道只能用于不同协程之间的通信
	// 测试3：注释掉 msg := <-messages，只有go func()里的发送方会如何？可以正常执行，也没报错。
	// 测试4：注释掉 msg := <-messages，并在主协程里(goroutine 1)往messages发信息 messages <- "ffff"。报和上面一样的deadlock错误
	// 结论: 只有send或receiver的确会阻塞, 但是这个阻塞只有在main函数执行里的时候才报上面的deadlock错误
	msg := <-messages

	fmt.Println(msg)

	// 👉 使用带缓冲区的channel，避免阻塞
	messagesBuffered := make(chan string, 2)

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	messagesBuffered <- "buffered"
	messagesBuffered <- "channel"

	// 使用带缓冲的channel就可以在自己的协程里收发
	fmt.Println(<-messagesBuffered)
	fmt.Println(<-messagesBuffered)

	// 无缓冲是同步的, 同一时刻，同时有 读、写两端把持 channel. 如果只有读端，没有写端，那么 “读端”阻塞, 如果只有写端，没有读端，那么 “写端”阻塞。
	// 有缓冲是异步的, 只有缓冲区被填满后，写端才会阻塞. 只有缓冲区被读空时，读端才会阻塞

	// 👉 遍历channel，range over channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // 这步如果不关闭的话，执行会报deadlock错误

	// it’s possible to close a non-empty channel but still have the remaining values be received.
	for ele := range queue {
		fmt.Println(ele)
	}
}
