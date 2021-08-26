package main

import "fmt"

func main() {
	// 使用make创建空的map, make初始化的时候可以指定len和cap
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("v1:", m["k1"])
	fmt.Println("len:", len(m))

	for k, v := range m {
		fmt.Println(k, v)
	}

	// 删除key
	delete(m, "k2")
	fmt.Println("len:", m)

	// 第二个返回true或false，表示key存不存在
	_, ok := m["k2"]
	if !ok {
		fmt.Println("k2 is not exsit")
	}

	// 声明的时候初始化map的内容
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// 下面的使用会包panic，使用map前需要先分配内存
	// var k map[string]string
	// k["k1"] = "v1"

	// fmt.Println(m == n) 👈 错误，map和slice一样是引用类型，只能和nil比较
}
