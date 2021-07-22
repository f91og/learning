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

func TryPointers() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针即地址，这里打印的应该时i的虚拟内存的地址
	fmt.Println("pointer:", &i)
}
