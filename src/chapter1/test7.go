package main

import (
	"bytes"
	"fmt"
)

func main() {
	//这里不能写成 b:=[]byte{"Golang"}  因为是强制类型装换
	b := []byte("Golang")
	subslice1 := []byte("go")
	subslice2 := []byte("Go")
	fmt.Println(bytes.Contains(b, subslice1))
	fmt.Println(bytes.Contains(b, subslice2))

	s2 := []byte("同学们，上午好")
	m := func(r rune) rune {
		if r == '上' {
			r = '下'
		}
		return r
	}
	fmt.Println(string(s2))
	// func Map(mapping func(r rune) rune, s byte[]) []byte
	//Map函数首先讲s转为UTF-8编码的字符序列
	//然后使用mapping将每个unicode字符映射为对应字符
	//最后将结果保存到一个新切片中
	fmt.Println(string(bytes.Map(m, s2)))

	s3 := []byte("google")
	old := []byte("o")
	new := []byte("ee")
	n := 1
	// func Replace(s, old, new []byte, n int) []byte
	// 返回一个切片的副本，并且将前n个不重叠的子切片old替换为new ,如果n<0则不限制替换数量
	fmt.Println(string(bytes.Replace(s3, old, new, n)))
	fmt.Println(string(bytes.Replace(s3, old, new, -1)))

	//将字节切片转化为对应的UTF-8编码的字节序列，并返回对应的Unicode切片
	s4 := []byte("中华人民共和国")
	r1 := bytes.Runes(s4)
	//func Runes(b []byte) []rune
	fmt.Println(string(s4), len(s4))
	fmt.Println(string(r1), len(r1))

	//字节切片的每个元素依旧是字节切片
	s5 := [][]byte{
		[]byte("你好"),
		[]byte("世界"), //这个逗号必不可少
	}
	sep := []byte(",")
	//func Join(s [][]byte, sep byte[]) byte[]
	// 用字节切片seq把s5中的每个切片连接成一个
	fmt.Println(string(bytes.Join(s5, sep)))

	//浮点数练习
	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	a := float64(20.0)
	c := 42
	fmt.Println(a)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", c)

}
