package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}
func main() {
	x := Vertex{3,4}
	fmt.Println(x)
	fmt.Println(Vertex{1,2})
}