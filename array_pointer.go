package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{100, 200, 300}
	var i int
	var pstr [MAX]*int

	for i = 0; i < MAX; i++ {
		pstr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *pstr[i])
		fmt.Printf("a[%d] = %d\n", i, pstr[i])
	}
}
