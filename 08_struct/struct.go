package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X: 3, Y: 4} // or just Vertex{1,2}
	p := &v
	p.X = 1e9 // 10^9

	v1 := Vertex{1, 2} // has type Vertex
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0
	c := &Vertex{1, 2} // has type *Vertex
	fmt.Println(v)
	fmt.Println(v1, c, v2, v3)
}
