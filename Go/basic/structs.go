package main

import "fmt"

// 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合
type person struct {
	name string // 注意结构体中的字段声明前面不需要var啥的，直接写变量名就可以了
	age  int
}

func newPerson(name string) *person {
	p := person{name: name} // 初始化一个结构体
	p.age = 14
	return &p
}

func TryStructs() {
	fmt.Println(person{"Bob", 16})
	fmt.Println(person{name: "Alice", age: 30})
	// fmt.Println(person{"Fred"}) 这种方式不可以
	fmt.Println(person{name: "Mark"})
	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}

	fmt.Println(newPerson("Jon")) // &{Jon 14}

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// the pointers are automatically dereferenced.
	// Dereferencing a pointer means getting the value that is stored in the memory location pointed by the pointer.
	// The operator * is used to do this, and is called the dereferencing operator.
	sp := &s
	fmt.Println(sp)  // &{Sean 50}, 注意这里没有自动dereference
	fmt.Println(*sp) // {Sean 50}，使用 * operator俩dereference
	sp.age = 51      // 这里的使用好像就和C++不一样了，
	fmt.Println(sp.age)
}
