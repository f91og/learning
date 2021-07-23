package main

import (
	"fmt"
	"time"
)

func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job")
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job")
		results <- j * 2
	}
}

// 下面的例子是用3个worker来处理5个job，首先循环起3个worker，它在读jobs通道的时候都会阻塞等待，然后在
// 主协程中往jobs通道里放入编号从1到5的job，让这3个worker去并行竞争处理这5个job，每个worker都将处理结果放入主协程指定的results通道中
func TryWorkPools() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results // Finally we collect all the results of the work, 有缓冲区的通道读时缓冲区为空的时候才会阻塞
	}

}
