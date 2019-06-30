package main

import (
	"fmt"
)

type Vertex2 struct {
	X, Y int
}

func main() {

	o := Vertex2{12, 124}
	defer fmt.Println(o)
	o.X = 43
	fmt.Println("o = ")
	fmt.Println(o)
}
