package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures
// 和结构体相反，接口是method签名的集合(注意是method，不是一般的function).
// 在go中，接口是合约，接口是一种类型

type geometry interface {
	area() float64
	perim() float64
}

// 同一个包中的methods.go中已经定义过了rect结构体，所以这里不能定义一样的
// type rect2 struct {
// 	width, height float64
// }

type rect2 struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect2) area() float64 {
	return r.width * r.height
}

func (r rect2) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 定义一个函数，参数利用接口来起到泛型和多态的作用，只要是geometry类型的接口都可以传入这个方法内
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func TryInterfaces() {
	r := rect2{width: 3, height: 4}
	c := circle{radius: 5}

	// rect和circle的方法签名和geometry接口中定义的一致，所以可以用这两个结构体的实例来传入
	// 空接口可以接受任何类型的值
	// 类型断言只能作用于接口类型的变量上（只有接口类型才有动态类型），这里传入r更像是一种合约的使用，r不能看成是接口类型。
	fmt.Println(r)
	measure(r)
	measure(c)
}
