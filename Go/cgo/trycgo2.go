package main

// 先自定义一个SayHello的C函数，然后再go里面调用

/*
#include <stdio.h>

static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"

func TryCgo3() {
	C.SayHello(C.CString("Hello, World\n"))
}
