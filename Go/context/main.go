package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// 上下文 context.Context Go 语言中用来设置截止日期、同步信号，传递请求相关值的结构体
// context.Context 是 Go 语言在 1.7 版本中引入标准库的接口1，该接口定义了四个需要实现的方法，其中包括：
// 1. Deadline() — 返回 context.Context 被取消的时间，也就是完成工作的截止日期
// 2. Done() — 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭，多次调用 Done 方法会返回同一个 Channel
// 3. Err() — 返回 context.Context 结束的原因，它只会在 Done 方法对应的 Channel 关闭时返回非空的值
// 4. Value() — 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据
func hello(w http.ResponseWriter, req *http.Request) {
	// A context.Context is created for each request by the net/http machinery, and is available with the Context() method
	ctx := req.Context()

	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n") // 把 hello\n 写入到 w 中
	case <-ctx.Done(): // While working, keep an eye on the context’s Done() channel for a signal that we should cancel the work and return as soon as possible.
		err := ctx.Err() // The context’s Err() method returns an error that explains why the Done() channel was closed.
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError) //
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

// https://www.sohamkamani.com/golang/context-cancellation-and-values/

// 👇Context Cancellation
// use <-ctx.Done() to listen a cancellation event
// for emitting a cancellation event, use WithCancel function in the context package, eg: ctx, fn := context.WithCancel(ctx)

// assume that this operation failed for some reason
func operation1(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done(): // operation2监听使用的context有没有被取消
		fmt.Println("halted operation2")
	}
}

func TryContextCancellation() {
	ctx := context.Background()

	// Create a new context, with its cancellation function from the original context
	// 利用WithCancel()可以得到一个cancel()方法，用这个方法就可以 emit a cancellation event
	ctx, cancel := context.WithCancel(ctx)

	// we want the context used in operation2 will be cancled if operation1 failed
	go func() {
		err := operation1(ctx)
		// If this operation returns an error cancel all operations using this context
		if err != nil {
			cancel()
		}
	}()

	operation2(ctx)
}

// 👇 Context Timeouts
// eg: ctx, cancel := context.WithTimeout(ctx, 3*time.Second), The context will be cancelled after 3 seconds
// eg: ctx, cancel := context.WithDeadline(ctx, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)), the context will be cancelled on 2009-11-10 23:00:00

// consider making an HTTP API call to an external service. If the service takes too long, it’s better to fail early and cancel the request
func TryContextTimeouts() {
	// Create new context With a deadline of 100 milliseconds
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 100*time.Millisecond)

	// make a http request
	req, _ := http.NewRequest(http.MethodGet, "http://Google.com", nil)
	// Associate the cancellable context we just created to the request
	req = req.WithContext(ctx)

	// Create a new HTTP client and execute the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("request failed", err)
		return
	}
	fmt.Println("Response received, status code:", res.StatusCode)
}

// 👇 Context Values
// the values that need to be transfered across methods ro goroutines, eg: an id
// Using the context variable to pass down operation-scoped information is useful for a number of reasons:
// 1. It is thread safe: You can’t modify the value of a context key once it has been set. The only way set another value for a given key is to create another context variable using context.WithValue
// 2. It is conventional: The context package is used throughout Go’s official libraries and applications to convey operation-scoped data. Other developers and libraries generally play nicely with this pattern
const keyID = "id"

func TryContextValues() {
	rand.Seed(time.Now().Unix())
	ctx := context.WithValue(context.Background(), keyID, rand.Int())
	operation3(ctx)
}

func operation3(ctx context.Context) {
	// do some work
	log.Println("operation1 for id:", ctx.Value(keyID), " completed")
	operation4(ctx)
}

func operation4(ctx context.Context) {
	// do some work
	// this way, the same ID is passed from one function call to the next
	log.Println("operation2 for id:", ctx.Value(keyID), " completed")
}
