package main

import (
	"fmt"
)

type Vertex3 struct {
	X, Y int
}

func main() {

	v := Vertex3{1, 2}

	p := &v

	fmt.Println(p)
	fmt.Println(*p)

	p.X = 1e9
	fmt.Println(v)
}
