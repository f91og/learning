package main

import (
	"fmt"
	"os"
)

// 一个defer的面试题
func deferInterview1() int {
	x := 5
	defer func() { x++ }()
	// 返回值是5，这里的x是值传递，所以就算defer中修改了x也对原来的返回值没啥影响
	// 即return的内部逻辑是把x先赋给返回值变量，再返回这个返回值变量，如果这里引用类型则会返回值会被影响
	return x
}

func deferInterview2() (x int) {
	defer func() { x++ }()
	return 5 // 返回值是6，这里确定了返回的变量就是x，所以执行逻辑先是将x赋值为5，然后defer再自增x，最后将x返回
}

func deferInterview3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值是5
}

func deferInterview4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 返回值是5
}

// panic/recover
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复
		// panic层层上报，可以在funcB的上层recover，如果到最上层都没有recover则程序会因为panic退出
		// recover必须和defer搭配使用，定义recover的defer一定要在panic之前
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
	// panic之后的语句执行不到
}

func funcC() {
	fmt.Println("func C")
}

func main() {
	// defer语句，在return赋值和RET指令之间执行一些逻辑，主要用于资源清理，文件关闭，解锁，记录时间等
	f := createFile("/defer.txt")
	defer closeFile(f)
	writeFile(f)

	fmt.Println(deferInterview1(), deferInterview2(), deferInterview3(), deferInterview4())

	funcA()
	funcB() // 如果在B中不revocer，则C无法执行，报了panic后就会退出
	funcC()
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
