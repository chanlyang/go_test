package main

import "fmt"

func main() {
	var num int = 10
	i := 4.0
	fmt.Println(Sqrt(i), num)
}

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 56; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

func Sqrt2(x float64) float64 {
	a := 0.0

	a = (x/a + a) / 2

	return a
}
