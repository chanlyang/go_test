package main

import (
	"fmt"
	"math"
)

/*
定义一个接口
*/
type shape interface {
	area() float64
}

/*
定义一个circle
*/
type circle struct {
	x, y, radius float64
}

/*
定一个rectangle
*/
type rectangle struct {
	width, height float64
}

//定一个circle方法 实现shape.area()
func (circle circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

//定义一个rectangle方法实现shape.area()
func (rectangle rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func getArea(shape shape) float64 {
	return shape.area()
}

func main() {
	circle := circle{x: 0, y: 0, radius: 5}
	rectangle := rectangle{2, 3}
	fmt.Printf("circle area is :%f\n", getArea(circle))
	fmt.Printf("rectangle area is :%f\n", getArea(rectangle))
}
