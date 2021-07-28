package main

import "fmt"

//
func TrySlices() {
	// 使用make函数来创建切片，猜测是因为创建时不能在[]填入数字来指定长度(否则就变成了数组)，所以才需要make函数
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	fmt.Println("sl1:", s[2:5])
	fmt.Println("sl2:", s[:5])
	fmt.Println("sl3:", s[2:])

	// 和数组不同的是，定义的时候不给长度就是切片
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	// 先定义，再往切片中放入值
	var k []string
	fmt.Println(append(k, "k"))

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
