package main

import "fmt"

type myInt int     // 自定义类型，直接使用int无法给其添加方法
type yourInt = int // 类型别名

// 可以为自定义的类型添加方法，只能给自己的包里的自定义类型添加方法
func (m *myInt) printAddr() {
	fmt.Println(m)
}

// 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合
// 结构体小写开头表示这个结构体不可以在其他包使用，和函数一样
type person struct {
	name string // 注意结构体中的字段声明前面不需要var啥的，直接写变量名就可以了
	age  myInt
}

func newPerson(name string) *person {
	p := person{name: name} // 初始化一个结构体
	p.age = 14
	return &p
}

// 为结构体定义方法，这个结构体叫做这个方法的接收者，接收者多用类型名首字母的小写，比如这里的p
func (p *person) Speak(s string) {
	fmt.Println(s)
}

// 注意方法接收者是值和指针时的不同，值的传递是拷贝传递，所以这个方法对age的修改无法反映到原结构体p
// 我猜这样的方法在底层的执行会变成 func(p, ....) 这种函数形式
func (p person) guonian() {
	p.age++
}

// 嵌入结构体，匿名，方法继承
type NewType struct {
	person // 这里直接写入类型，没有给变量名，可以认为这里的person就是父类，因为NewType接收了person的所有方法
	nt     string
}

func (q *NewType) Strike(s string) {
	fmt.Println("Strike:", s)
}

// 嵌入结构体对接口方法签名的测试
type Speaker interface {
	Speak(s string)
}

func main() {
	fmt.Println(person{"Bob", 16})
	fmt.Println(person{name: "Alice", age: 30})
	// fmt.Println(person{"Fred"}) 这种方式不可以
	fmt.Println(person{name: "Mark"})
	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}

	fmt.Println(newPerson("Jon")) // &{Jon 14}

	s := person{name: "Sean", age: 50}
	fmt.Printf("%T\n", s.age) // main.myInt
	s.age.printAddr()         // 0xc0000045b0
	var m yourInt = 100
	fmt.Printf("%T\n", m) // int

	// the pointers are automatically dereferenced.
	// Dereferencing a pointer means getting the value that is stored in the memory location pointed by the pointer.
	// The operator * is used to do this, and is called the dereferencing operator.
	sp := &s
	fmt.Println(sp)  // &{Sean 50}, 注意这里没有自动dereference
	fmt.Println(*sp) // {Sean 50}，使用 * operator俩dereference
	sp.age = 51
	fmt.Println(sp.age)

	// 使用为结构体定义的方法
	s.Speak("hello world!")

	// 嵌入结构体，匿名，方法继承
	n := NewType{s, "newType"}
	fmt.Println(n.name, n.age)       // 实际上是 person.name, person.age
	n.Speak("实际上调用的是person的Speak方法") // 嵌入结构体的方法查找会层层向下，并且在同级层不可以有重名的方法

	// 嵌入结构体对接口方法签名的测试
	useSpeakerInterface(&s, "person say hello")
	useSpeakerInterface(&n, "new type say hello") // 被嵌入的结构体也能传入到这个函数中，不过执行的时候还是里层的person执行的Speak方法

	// 匿名结构体，多用于临时场景
	var ns struct {
		x string
		y string
	}
	ns.x, ns.y = "x", "y"

	fmt.Println("年前age：", s.age) // 51
	s.guonian()
	fmt.Println("年后age：", s.age) // 51

	n.guonian() // 嵌入的结构体的对外部包不可见的方法也会被被嵌入的结构体接收
}

// 这个函数接受的参数是实现了接口Speaker里的所有方法签名的对象
func useSpeakerInterface(i Speaker, c string) {
	i.Speak(c)
}
