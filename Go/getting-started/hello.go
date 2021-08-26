// 定义包名，包名和具体的文件路径无关, 和文件名也没有关系
// package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包
//  同一个文件夹下的文件只能有一个包名，否则编译报错
package main

import (
	"fmt"

	"rsc.io/quote"
)

// main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）
func main() {
	fmt.Println(quote.Go())
}
