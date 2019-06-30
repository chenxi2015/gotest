package main

import (
	"fmt"
)

func main() {

	fmt.Println("counting")

	count := 10
	for i := 0; i < count; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
