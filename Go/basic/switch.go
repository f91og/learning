package main

import (
	"fmt"
	"time"
)

func TrySwitch() {
	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend")
	default:
		fmt.Println("It is a weekday")
	}

	t := time.Now()
	// switch without an expression is an alternate way to express if/else logic
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}

	// 在go中接口是一种数据类型
	// interface{}, all the types implement the empty interface, 空接口可以保存任何值，也可以从空接口中取出原值
	// 静态类型：量声明的时候的类型，如var age int
	// 动态类型：程序运行时系统才能看见的类型
	// 因为空接口可以接受任何值，var i interface{}，i = 18，i = "Go编程时光" 。i 的静态类型就是 interface{}，但是它的动态类型先是int然后变成了string
	// 可以将任何类型看成空接口的实现。动态类型只是对于接口而言，即对于var i int = 100这样一开始定下了类型的变量，没有动态类型
	// https://segmentfault.com/a/1190000022931452
	whatAmI := func(i interface{}) {
		switch t := i.(type) { // type switch，switch和type Assertions组合起来的go的固定用法。本来的type Assertions是用于提取接口的值
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("i am a int")
		default:
			fmt.Printf("Don't know type %T\n", t) // %T-类型，v%-值
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
