package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := struct {
		a byte
		b byte
		c byte
		d int64
	}{0, 0, 0, 0}

	//将结构体指针转化为通用指针
	p := unsafe.Pointer(&s)
	//保存结构体地址备用 偏移量为0
	up0 := uintptr(p)
	//将通用指针类型转为byte指针类型
	pb := (*byte)(p)
	//给转换后的指针赋值
	*pb = 40
	//结构体内容跟着改变
	fmt.Println(s)

	//偏移到第二个字段
	up := up0 + unsafe.Offsetof(s.b)
	//将偏移后的地址转为通用指针
	p = unsafe.Pointer(up)
	//将通用指针转为byte类型指针
	pb = (*byte)(p)
	//给转换后的指针赋值
	*pb = 50
	//结构体改变
	fmt.Println(s)

	//偏移到第三个字段
	up = up0 + unsafe.Offsetof(s.c)
	//将偏移后的地址转为通用指针
	p = unsafe.Pointer(up)
	//将通用指针转为byte类型指针
	pb = (*byte)(p)
	//给转换后的指针赋值
	*pb = 60
	//结构体跟着变化
	fmt.Println(s)

	//偏移到第4个字段
	up = up0 + unsafe.Offsetof(s.d)
	p = unsafe.Pointer(up)
	pi := (*int64)(p)
	*pi = 70
	fmt.Println(s)
}
