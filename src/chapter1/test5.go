package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "process mode")

func main() {
	var n = 10
	ptr := &n
	fmt.Println(ptr)
	fmt.Println(*ptr)
	//%T 按go的格式输出
	fmt.Printf("%T\n", ptr)

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

	//使用指针输出命令行参数
	//解析命令行参数
	flag.Parse()
	fmt.Println(*mode)

	//定义指针
	str := new(string)
	*str = "chenliang 你好"
	fmt.Println(*str)
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func swap2(a, b *int) {
	*a, *b = *b, *a
}
