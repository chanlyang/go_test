package main

import (
	"fmt"
	"math"
)

//定义一个全局变量

var a int

func main() {

	var a int = 2
	var b int = 3
	fmt.Println("第一个变量:", a)
	fmt.Println("第二个变量:", b)
	result := sum(a, b)
	fmt.Println("返回的结果:", result)

	fmt.Printf("%3f\n", math.Pi)

}

func sum(a, b int) int {
	fmt.Println("第一个参数:", a)
	fmt.Println("第二个参数:", b)
	num := a + b
	return num
}
