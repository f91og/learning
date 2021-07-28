package main

import "fmt"

func TryMaps() {
	// 使用make创建空的map, make初始化的时候可以指定len和cap
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	// 删除键的方法和python一样
	delete(m, "k2")
	fmt.Println("len:", m)

	// 第二个返回true或false，表示key存不存在
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 声明的时候初始化map的内容
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// 下面的使用会包panic
	// var k map[string]string
	// k["k1"] = "v1"
}
