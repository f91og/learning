package main

import (
	"fmt"
	"sort"
)

func main() {
	// 使用make函数来创建切片，从数组创建的切片没办法指定len和cap，所以需要make函数
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0], s[1], s[2] = "a", "b", "c"
	fmt.Println("s[2]:", s[2], "s len:", len(s), "cap:", cap(s))

	s = append(s, "d")
	s = append(s, "e", "f", "g")                // 扩容会导致底层数组的改变（换了一个）
	fmt.Println("apd:", len(s), "cap:", cap(s)) // apd: 7 cap: 12,

	s = append(s, []string{"h", "i", "q"}...) // 利用...拆箱
	fmt.Println(s)

	c := make([]string, len(s)) // cppy前需要先开辟内存空间，否则copy不了
	copy(c, s)
	s[0] = "s"
	fmt.Println("c:", c) // copy函数会将底层数组拷贝，从而避免对原切片里数组的修改影响到新的切片

	fmt.Println("sl1:", s[2:5], "sl2:", s[:5], "sl3:", s[2:])

	// 和数组不同的是，定义的时候不给长度就是切片
	t := []string{"g", "h", "i"}
	fmt.Println("t len:", len(t), "t cap:", cap(t))

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

	// 从数组创建切片
	a1 := [...]int{1, 2, 3, 4, 5}
	s3 := a1[0:2]                                       // 数组切割后得到的是切片
	fmt.Println("s3 len:", len(s3), "s3 cap:", cap(s3)) // s3 len: 2 s3 cap: 5
	s4 := a1[1:2]
	fmt.Println("s4 len:", len(s4), "s3 cap:", cap(s4)) // s4 len: 1 s3 cap: 4, 容量从开始往后计算

	a1[0] = 45
	fmt.Println(s3) // [45 2], 从数组创建的切片的底层的数组就是原数组

	// fmt.Println(s3 == s4)👈错误，切片只能和nil比较
	// 应该利用len(s)==0来判断是否是空切片

	// 切片元素的删除，Go中没有删除切片元素的专门方法，需要自己实现
	// a1 = append(a1[:2], a1[3:]...) 👈 错误，a1是数组，用不了append方法
	s5 := a1[:] // 直接将数组等价转换为切片
	s5 = append(s5[:2], s5[3:]...)
	fmt.Printf("%v\n", s5) // [45 2 4 5]

	sort.Ints(s5)
	fmt.Println(s5) // [2 4 5 45]，对切片排序

}
