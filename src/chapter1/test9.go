package main

import "fmt"

func main() {
	//声明一个长度为5的整数数组
	//数组一旦声明 它的数据类型和长度就不能再改变
	var array1 [5]int

	fmt.Printf("array1: %d\n", array1)
	array1[1] = 1

	//声明一个长度为5的整数数组并初始化每个元素
	array2 := [5]int{12, 13, 14, 15, 16}
	array2[1] = 19
	fmt.Printf("array2: %d\n", array2)

	//n是一个长度为10的数组
	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 11
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("element[%d] = %d\n", j, n[j])
	}

	//二维数组
	var arr = [5][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}
	var e, f int
	for e = 0; e < 5; e++ {
		for f = 0; f < 2; f++ {
			fmt.Printf("a[%d][%d] = %d\n", e, f, arr[e][f])
		}
	}

	//可以根据初始化数据的个数设置数组大学
	var array3 = [...]float64{12.12, 32.23, 213.432}
	fmt.Println(array3)

}
