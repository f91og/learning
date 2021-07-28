package main

import (
	"fmt"
	"time"
)

func TryTimersAndTickers() {
	// Timers represent a single event in the future.
	// You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
	// C是timer里自带的channel
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// If you just wanted to wait, you could have used time.Sleep. One reason a timer may be useful is that you can cancel the timer before it fires. Here’s an example of that.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop() //
	if stop2 {
		fmt.Println("Timer2 stopped")
	}
	time.Sleep(2 * time.Second)

	// Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular interval
	ticker := time.NewTicker((500 * time.Millisecond))
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
