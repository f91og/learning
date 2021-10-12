package main

import (
	"fmt"
	"math"
)

// 接口是一种类型，内部的存储原理是存储了 动态类型和动态值
// 动态类型用于描述运行时接收的实现了接口方法的不同类型的struct(或者自定义类型)
// 动态值用于指向具体存储的那个对象，实际调用方法的时候是动态值里指向的对象调用自己的方法
type geometry interface {
	area() float64
}

// 结构嵌套
type advancedGeometry interface {
	geometry
	advancedCalculation() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// 定义一个函数，参数利用接口来起到泛型和多态的作用，只要是geometry类型的接口都可以传入这个方法内
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

// 实现接口的方法时，指针作为接收者和值作为接收者时的区别
type square struct {
	edge float64
}

func (s *square) area() float64 {
	return s.edge * s.edge
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	s := square{6}

	// rect和circle的方法签名和geometry接口中定义的一致，所以可以用这两个结构体的实例来传入
	measure(r)
	measure(c)

	var d geometry              // 接口是一种类型，这里定义接口变量时，不需要考虑是不是指针
	fmt.Printf("%T，%v\n", d, d) // nil, nil 这里打印的是接口里存的动态类型和动态值，一开始啥都没有，所以是nil

	// 当用指针作为方法实现的接收者时，实现了接口的是这个对象的指针而不是这个对象的值
	// d = s // 编译不通过
	d = &s
	fmt.Println(d.area()) // 36

	d = r
	fmt.Printf("%T，%v\n", d, d) // main.rect2，{3 4}
	d = c
	fmt.Printf("%T，%v\n", d, d) // main.circle，{5}

	// 当用值作为方法的接收者时，这个对象的值和指针都实现了这个方法对应的接口
	d = &c
	fmt.Printf("%T，%v\n", d, d) // *main.circle，&{5}

	// 空接口可以接收任何值
	var any interface{} = 123
	fmt.Println(any) // 123
	any = "any"
	fmt.Println(any)

	// 类型断言只能作用于接口类型的变量上（只有接口变量才存储了对象的动态类型）
	// 类型断言1，取接口里存储的动态值
	fmt.Println(d.(geometry), d.(*circle)) // &{5} &{5}
	// fmt.Println(d.(rect2)) 👈 panic
	// fmt.Println(any.(int)) 👈 panic

	// 类型断言2，判断接口里存储的动态类型
	if v, ok := any.(int); ok {
		fmt.Println(v)
	} else if v, ok := any.(string); ok {
		fmt.Println(v)
	}

	// 类型断言2，判断接口里存储的动态类型，.(type) 只能和switch一起用
	switch t := any.(type) {
	case string:
		fmt.Printf("any is a string，value is %v", t)
	case int:
		fmt.Printf("any is a int, value is %v", t)
	default:
		fmt.Printf("unsupport type")
	}

}
