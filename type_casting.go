package main

import "fmt"

func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Println("mean的值是：", mean, "\n")
	fmt.Printf("mean的值是：%f\n", mean)
}
