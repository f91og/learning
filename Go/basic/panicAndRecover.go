package main

import "os"

// panic和recover关键字，panic用于抛出问题并终止后面执行，recover用于再defer中处理这个问题，若不处理则函数调用栈会层层向上panic
// panic和recover是一对，是内置函数, defer是关键字
// A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle
func TryPanicAndRecover() {

	// panic会停止当前goroutine的正常执行，当函数F调用panic时，F的正常执行被立即停止，然后运行所有在F函数中的defer函数
	// 如果在defer中没有recover这个panic则对于调用F的函数G，F就相当于panic，终止G的执行并运行G中所defer函数，依次类推，如果调用栈中的函数都没有recover，则程序异常退出
	// 如果F的defer中有recover捕获，则F在执行完defer后正常返回，调用函数F的函数G继续正常执行
	// recover都是在当前的goroutine里进行捕获的，对于创建goroutine的外层函数，如果goroutine内部发生panic并且内部没有用recover，外层函数是无法用recover来捕获的
	// defer 定义的函数如果在 panic 后面，该函数在 panic 后就无法被执行到
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
