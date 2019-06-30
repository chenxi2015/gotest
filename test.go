package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println("hello world")
	r := RandInt(10, 13)
	fmt.Println(r, "time:", time.Now().UnixNano())
	t := time.Now().UnixNano()
	sum := 0
	for i := 0; i <= 40000+r; i++ {
		for j := 0; j <= 40000; j++ {
			sum = sum + i*j
		}
	}
	i := float64(time.Now().UnixNano()-t) / (1000000 * 1000)
	fmt.Printf("sum: %d", sum)
	fmt.Println("")
	fmt.Println("second:", i)
}

func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
