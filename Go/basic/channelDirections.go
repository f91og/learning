package main

import "fmt"

// 通道作为函数参数时，可以指定其方向，从而限定通道在函数中使用时只能发或者只能收
// specify if a channel is meant to only send or receive values when using channels as function parameters
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func TryChannelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// 实际传参的时候还是将通道整体传入，只不过通过函数参数列表中规定的通道方向，可以限定这个通道在函数中使用时只能发或者只能收
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
