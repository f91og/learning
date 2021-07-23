package main

import (
	"fmt"
	"time"
)

// Go supports rate limiting with goroutines, channels, and tickers.

// 利用通道的阻塞特性和tickers的定时执行来实现可控制的速度控制器
func TryRateLimiting() {
	// 假设我们想限制对到来的request的处理速度
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 利用通道的阻塞机制+tick的定时执行来使对request的处理按一定时间间隔间隔化
	// This limiter channel will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the limiter channel before serving each request, we limit ourselves to 1 request every 200 milliseconds
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit
	// We can accomplish this by buffering our limiter channel. This burstyLimiter channel will allow bursts of up to 3 events
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting，先放3个，这样一开始的rquests处理不会被时间间隔隔开，能够被全力处理
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		// Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
		// 这里先会被阻塞，因为上面已经实现将3个time.Now()放入到burstyLimiter了
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming requests. The first 3 of these will benefit from the burst capability of burstyLimiter
	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)
	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
