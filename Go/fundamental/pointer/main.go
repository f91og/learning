package main

import "fmt"

// 这里按值传递参数，zeroval函数里会得到ival的一份拷贝，对它的修改不会反映到原本的参数上
func zeroval(ival int) {
	ival = 0
}

// 指针，按照地址引用传递，传入的时iptr的地址，函数内的修改会反映到原本的参数上
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针是一个存储在内存里的变量，有自己的内存地址，只不过它的值是别的变量的内存地址
	fmt.Println("pointer:", &i)

	// & 取址，* 取指针所指的内存的内容
	j := 10
	p := &j
	fmt.Printf("%T, %v\n", p, p) // *int, 0xc0000120b8(变量i的内存地址)
	fmt.Println(*p)

	// var q *int = &j

	// var q *int
	// *q = 100
	// fmt.Println(q) // panic

	q := new(int) // 使用new开辟内存并返回指针，除此之外可以在不需要声明一个变量的情况下使用非空指针
	fmt.Println(*q)

	fmt.Println(p == q) // false
	q = &j
	fmt.Println(p == q) // true，指向同一个地址的变量时，指针的比较结果才是true
}
