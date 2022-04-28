package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	var target int = 3
	fmt.Println(towSum(arr, target))

}

func towSum(arr []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		another := target - arr[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[arr[i]] = i
	}
	return nil
}
