package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// 利用sync.WaitGroup来实现等待一组goroutine执行完毕
// a WaitGroup must be passed to functions by pointer
func worker3(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Done()方法将WaitGroup的counter计数器减1

	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	// 这里是如何实例化一个WaitGroup的？貌似直接声明一个就开始用了？👈 因为WaitGroup是结构体，结构体不需要像channel那样先make分配内存再使用，
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker3(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
	wg.Wait()

	// 互斥锁，sync.Mutex
	var mutex = &sync.Mutex{}
	go addWithMutex(mutex)
	go addWithMutex(mutex)
	// 读写互斥锁，sync.RWMutex，适合读多写少的场景

	// 利用sync.Once来实现并发安全的只执行一次
	for i := 0; i < 10; i++ {
		go func() {
			_ = ReadConfig() // 只会打印一次 init config
		}()
	}
	time.Sleep(time.Second)

	// 利用内置的并发安全的sync.Map，开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用
	var m = sync.Map{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()

	// sync.Pool，对象池
	// var studentPool = sync.Pool{
	// 	New: func() interface{} {
	// 		return new(Student)
	// 	},
	// }
	// stu := studentPool.Get().(*Student)	// 返回值是 interface{}，因此需要类型转换
	// json.Unmarshal(buf, stu)
	// studentPool.Put(stu)	// 对象使用完毕后，返回对象池

	// sync.Cond
}

// sync.Once，确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等
// 使用 sync.Once 读取环境变量并转为对应的配置
type Config struct {
	Server string
	Port   int64
}

var (
	once   sync.Once
	config *Config
)

func ReadConfig() *Config {
	once.Do(func() {
		var err error
		config = &Config{Server: os.Getenv("TT_SERVER_URL")}
		config.Port, err = strconv.ParseInt(os.Getenv("TT_PORT"), 10, 0)
		if err != nil {
			config.Port = 8080 // default port
		}
		log.Println("init config")
	})
	return config
}

// 互斥锁
var x int64

func addWithMutex(m *sync.Mutex) {
	for i := 0; i < 5000; i++ {
		m.Lock()
		x = x + 1
		m.Unlock()
	}
}
