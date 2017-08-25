package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
	
}

func main()  {
	
	o := Vertex{12,124}
	defer fmt.Println(o)
	o.X = 43
	fmt.Println(o)
}