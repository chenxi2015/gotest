package main

import (
	"fmt"
)

type Vertex1 struct {
	X int
	Y int
}

func main() {
	x := Vertex1{3, 4}
	fmt.Println(x)
	fmt.Println(Vertex1{1, 2})
}
