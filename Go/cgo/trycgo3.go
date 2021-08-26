package main

// _GoString_预定义的C语言类型，用来表示Go语言字符串，目的是为了直接使用Go字符串

//void SayHello2(_GoString_ s);
import "C"
import "fmt"

func TryCgo4() {
	C.SayHello2("Hello, World\n")
}

// 通过CGO的//export SayHello指令将Go语言实现的函数SayHello导出为C语言函数

//export SayHello2
func SayHello2(s string) {
	fmt.Print(s)
}
