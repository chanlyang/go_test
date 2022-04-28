package main

import (
	"fmt"
	"sort"
)

func main() {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}
	fmt.Println(megerTowArray(num1, num2))
	fmt.Println(megerTowArrayBySort(num1, num2))
}

//两个数组，找出公共的元素
func megerTowArray(num1 []int, num2 []int) []int {
	//使用map先把第一个数组的元素作为key,个数作为value放入map中
	//遍历第二个数组，如果元素在map中，则将该元素放到
	m := make(map[int]int)
	for _, v := range num1 {
		m[v] += 1
	}
	k := 0
	for _, v := range num2 {
		if m[v] > 0 {
			m[v] -= 1
			num2[k] = v
			k++
		}
	}
	return num2[0:k]
}

//如果是排好序的两个数组呢
func megerTowArrayBySort(num1 []int, num2 []int) []int {
	sort.Ints(num1)
	sort.Ints(num2)
	i, j, k := 0, 0, 0
	for i < len(num1) && j < len(num2) {
		if num1[i] > num2[j] {
			j++
		} else if num1[i] < num2[j] {
			i++
		} else {
			num2[k] = num1[i]
			i++
			j++
			k++
		}
	}
	return num2[0:k]
}
