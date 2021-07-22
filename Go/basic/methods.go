package main

import "fmt"

type rect struct {
	width, height int // 一行声明多个变量
}

// 定义方法，其实就是一个函数。定义的格式为：func (挂载在哪个结构体上使用) 方法名(参数列表) 返回值列表 {...方法的具体逻辑...}
// 在方法的具体逻辑中可以使用方法名前面指定挂载的结构体

// 为rect定义一个计算它面积的方法
func (ptr *rect) area() int {
	return ptr.width * ptr.height
}

// Methods can be defined for either pointer or value receiver types

// 为rect定义一个计算它周长的方法
func (ptr rect) perim() int {
	return 2 * (ptr.width + ptr.height)
}

func TryMethods() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for method calls
	// You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
