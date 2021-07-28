package main

import "fmt"

// 闭包就是封装了一些不是它自己的东西的函数

// 下面这个函数放回另一个返回值是int的函数, 这个函数把一些数据(i)和对他的操作函数都封装在一起了，其实和这个函数自己没啥关系，
// 只是将数据和函数封装在一起，从功能的角度来说和一般的函数职责不一样，因此被称为闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TryClosures() {
	nextInt := intSeq() // 这里nextInt是一个函数

	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
