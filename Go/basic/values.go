package main

// 一个包下只能有一个mian函数
// func main() {

// }

import "fmt"

func TryValues() {
	fmt.Println("go" + "lang") // 字符串拼接
	fmt.Println("1+1=", 1+1)
	fmt.Println("7/3", 7/3)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
