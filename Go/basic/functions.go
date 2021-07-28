package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// return multiple value
func vals() (int, int) {
	return 3, 7
}

// 接受可变参数的函数，这个函数没有返回值，所以定义函数时最后面指定返回值类型的部分可以省略
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func TryFunctions() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)

	a, b := vals()
	fmt.Println("a,b:", a, b)

	// _, c := vals()
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
