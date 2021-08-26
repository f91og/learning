package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(s[1:]) // ello, world

	fmt.Printf("%#v\n", []byte("Hello, 世界")) // []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c}
	fmt.Println("\x48")                      // H
	fmt.Println("\xe4\xb8\x96")              // 世

	fmt.Printf("%#v\n", []rune("世界"))             // []int32{19990, 30028}
	fmt.Printf("%#v\n", string([]rune{'世', '界'})) // 世界
}
