package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	fmt.Println(Vertex{1, 2})
	//结构体字段使用点号来访问
	v := Vertex{3, 4}
	v.X = 5
	fmt.Println(v)

	//结构体字段可以通过结构体指针来访问
	e := Vertex{5, 6}
	p := &e
	p.X = 7
	fmt.Println(e)

	var (
		v1 = Vertex{2, 3}  //类型为Vertex
		v2 = Vertex{X: 3}  //Y值被省略，Y:0
		v3 = Vertex{}      // X:0,Y:0
		pi = &Vertex{3, 4} //类型为 *Vertex, &返回一个指向结构体的指针
	)

	fmt.Println(v1, v2, v3, pi)

}
