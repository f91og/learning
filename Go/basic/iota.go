package main

import "fmt"

// go的iota关键字以及利用它实现枚举
func TryEnumWithIota() {
	// iota：特殊常量，每当 iota 在新的一行被使用时，它的值都会自动加 1
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i) // 0 1 2 ha ha 100 100 7 8

	// 利用iota来实现枚举
}
