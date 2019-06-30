package main

import (
	"fmt"
)

type St struct {
	X int
	Y int
}

var (
	v1 = St{1, 2}
	v2 = St{X: 1}
	v3 = St{}

	p = &St{1, 2}
)

func main() {

	fmt.Println(v1, v2, v3, p, *p)

}
