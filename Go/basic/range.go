package main

import "fmt"

func TryRange() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) // 使用Printf来使用格式化字符串，Println会照原样输出
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself.
	// 下面输出 0 103, 1 111
	// rune代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
