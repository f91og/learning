#include <stdio.h>

// 为了允许外部引用，所以需要去掉函数的static修饰符。关于C的staic关键字作用: 
// 1. 隐藏，未加 static 前缀的全局变量和函数都具有全局可见性
// 2. 保持变量内容的持久
// 3. 默认初始化为 0
void SayHello(const char* s) {
    puts(s);
}