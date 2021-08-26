package main

// 包含C语言的<stdio.h>头文件, 注释不能紧跟import "C"，否则会被当成C的执行逻辑

//#include <stdio.h>
import "C"

func TryCgo1() {
	println("hello cgo") // 直接使用C的println函数
}

func TryCgo2() {
	// 通过CGO包的C.CString函数将Go语言字符串转为C语言字符串
	C.puts(C.CString("Hello, World\n")) // 使用go的C包提供的方法
}
