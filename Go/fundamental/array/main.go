package main

// 存取速度：L1 cache > L2 > L3 > 内存
// 对数组的存储应该尽量直接通过L1缓存，这样有最大的存取效率
// 处理器用Prefetcher来提前预测需要用到的数据，将其预先装载到L1缓存
// 虽然程序对数据的处理粒度是以1字节为单位，但是从硬件缓存的角度来说并不是1字节为单位，而是以64个字节为的cache line为单位
// 程序需要创建一种可以预测存取访问的数据结构给Prefetcher，比如分配一块连续的内存空间。即数组/切片
// 数组分配之后，无论它大小是多少，每个元素与其他元素的距离都是相等的
// 矩阵的行遍历有最高的效率，因为Prefetcher可以根据数组内存空间连续的特点将其通过一个一个cache line预装载到L1缓存中，按列
// 遍历效率最差，似乎是一个随机存取的模式，会发生cache miss和TLB cache miss。LinkedList Traverse性能中等的原因是因为有Translation lookaside buffer
// caching system一次移动64个字节（根据不同机器有所不同？），但是操作系统按照每页4K来管理内存，操作系统给运行的程序提供地址空间相互独立的虚地址，然后通过页表机制映射到实际的物理内存地址上
// TLB: 内存里有页表，先查页表再去查内存中相应的地址。为了减少内存IO次数，将页表缓存在TLB中，专门用于改进虚拟地址到物理地址转换速度的缓存

import (
	"fmt"
	"reflect"
)

func main() {

	var a [5]int
	fmt.Println(a) // 和Python一样，int数组的默认值是0，string数组的默认值是""

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	fmt.Println(len(a)) // len()内置函数，可以查看数组或者切片的长度
	fmt.Println(cap(a)) // cap()内置函数, 可以查看数组或者切片的容量，数组定长所以这里结果一样

	b := [5]int{1, 2, 3, 4, 5} // 使用:=在定义的时候并初始化值
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	// 声明数组时只指定特定索引上的值，其他的赋初始值
	months := [...]string{1: "January", 12: "December"}
	fmt.Println(months[0], months[1], months[10])

	// 数组是值语义，当一个数组变量被赋值或者被传递的时候，实际上会复制整个数组。如果数组较大的话，数组的赋值也会有较大的开销。
	// 为了避免复制数组带来的开销，可以传递一个指向数组的指针
	var c = [...]int{1, 2, 3}
	var d = &c // d 是指向数组的指针
	fmt.Println(c[0], c[1])
	fmt.Println(d[0], d[1]) // 通过数组指针访问数组元素的方式和数组类似
	fmt.Println(d)          // &[1, 2, 3]
	fmt.Println(*d)         // * 取指针d指向的内存地址的值，[1 2 3]
	fmt.Println(&d)         // 0xc000006030, 指针d的地址，d是一个在内存中的变量，有地址和内容两个属性，地址是自己的地址，内容是数组d的内存地址

	for i, v := range d { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}
	for i := range d {
		fmt.Printf("d[%d]: %d\n", i, d[i])
	}

	// 用fmt.Printf函数提供的%T或%#v谓词语法来打印数组的类型和详细信息
	fmt.Printf("c: %T\n", c)  // c: [3]int
	fmt.Printf("c: %#v\n", c) // c: [3]int{1, 2, 3}
	fmt.Printf("d: %T\n", d)  // d: *[3]int
	fmt.Printf("d: %#v\n", d) // d: &[3]int{1, 2, 3}

	// 长度为0的数组在内存中并不占用空间，空数组虽然很少直接使用，但是可以用于强调某种特有类型的操作时避免分配额外的内存空间
	// 下面并不关心管道中传输数据的真实类型，其中管道接收和发送操作只是用于消息的同步。对于这种场景，可以用空数组（或者空结构体）来作为管道类型可以减少管道元素赋值时的开销
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1

	// 数组的数据结构
	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)
	fmt.Println(t)	// [3]int
	fmt.Println(v)	// [1 2 3]
}
