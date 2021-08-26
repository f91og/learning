package main

import (
	"fmt"
	"runtime"
	"time"
)

func f(s string) {
	fmt.Println("i am goroutine:", s)
}

func main() {
	// runtime.GOMAXPROCS(1) // 设置当前程序并发时占用的CPU逻辑核心数
	fmt.Println("当前机器拥有的CPU核心数是：", runtime.NumCPU())

	// 执行函数的时候在前面加上go让其在新协程中执行
	go f("goroutine 1")

	// go中函数里面可以定义函数，这里不光定义了函数，还定义后立马就执行
	go func(s string) {
		fmt.Println("i am goroutine:", s)
	}("goroutine 2")

	time.Sleep(time.Second)
	fmt.Println("done")

	// 使用通道来进行协程之间的通信，通道在协程执行完后会自动关闭
	messages := make(chan string) // 无缓冲的通道，在发送方和接收方都准备好之前会一直阻塞
	go func() {
		messages <- "ping"
		// 这句有可能打印不出来，因为main协程执行完程序就退出了，此时这个协程可能还没有执行完毕
		fmt.Println("send message goroutine done!")
	}()
	time.Sleep(1000)
	msg := <-messages
	fmt.Println("receieved message:", msg)

	// 👉 使用带缓冲区的channel，避免阻塞
	messagesBuffered := make(chan string, 2)
	messagesBuffered <- "buffered"
	messagesBuffered <- "channel"

	// 使用带缓冲的channel就可以在自己的协程里收发
	fmt.Println(<-messagesBuffered, <-messagesBuffered)

	// 读空的缓冲区和往缓冲区满的通道里写都会阻塞，这里阻塞里会造成死锁，所以会报dead lock error
	// fmt.Println(<-messagesBuffered)

	// 不会报dead lock error，因为mian协程没有被阻塞，即使下面这个协程被阻塞，也能够在main协程退出时结束
	go func() { fmt.Println(<-messagesBuffered) }()

	// 利用 range 遍历通道中的内容
	messagesBuffered <- "buffered1"

	// 下面这样直接range会报deadlock error，因为messagesBuffered没有关闭，range就会一直读(会取出通道里的内容)，读到没有的时候main协程就会一直阻塞在这里
	// for ele := range messagesBuffered {
	// 	fmt.Println(ele)
	// }

	// 关闭messagesBuffered后再用range去读，range会自动读到通道中没有值时为止
	close(messagesBuffered)
	for ele := range messagesBuffered {
		fmt.Println(ele) // buffered1
	}

	// 读已关闭的通道
	v, ok := <-messagesBuffered
	fmt.Println("读已关闭的通道的结果是：", v, ok) // "", false
	// 写已关闭的通道，Panic
	// messagesBuffered <- "buffered2"

	// 读写nil通道都会阻塞，关闭nil通道和已经关闭的通道都会引发panic

	pings, pongs := make(chan string, 1), make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// 通过go的select来等待多个channel操作
	c1, c2 := make(chan string), make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	// select会同时等待所有case，如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行，所以这里循环了两次，因为需要等待两个通道返回结果
	// 和switch一样可以用default语句。在没有case可以执行的时候，如果有defaul语句，那么select会执行default语句，否则select会一直阻塞，直到某个通信可以运行
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 通过worker pool来限制gorutine的数量
	TryWorkPools()
}

// 单向通道，限制通道在函数中只能发送或只能接收
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// worker pool，限制gorutine的数量
func workers(id int, jobs <-chan int, results chan<- int) {
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
		go workers(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results // Finally we collect all the results of the work, 有缓冲区的通道读时缓冲区为空的时候才会阻塞
	}
}
