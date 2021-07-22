package main

import "fmt"

var g int

func TryVariables() {
	var a = "initial"
	fmt.Println(a)

	// 可以省略int，go会自动推断类型
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// var e，不可以写成这种声明形式
	var e int
	fmt.Println(e) // 默认值0

	g = 100
	fmt.Println(g)

	// 这种方式只能在函数里使用
	f := "apple"
	fmt.Println(f)
}
