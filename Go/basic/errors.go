package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have type error, a built-in interface.
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message.
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// use custom types as errors by implementing the Error() method on them
type argError struct {
	arg  int
	prob string
}

//
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob) // 这里需要返回结果字符串，而不是直接输出到控制台，所以使用Sprintf方法
}

// 因为argError里实现了Error()这个error接口的方法签名，所以可作为error可以接收argError
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func TryErrors() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil { // if 和 for 一样，判断条件前面都可以加个语句
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 通过type assertion来取值, 这里go不会自动dereference
	// type assertion的两种用法：✌a.(int)这样取值. ✌switch和a.(type).
	// type assertion一定要作用于接口类型之上。原本的结构体上使用(x)，将其用一个接口类型接收后使用(✔)
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
