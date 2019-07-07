package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum发送到通道c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // 17
	go sum(s[len(s)/2:], c) // -5

	x, y := <-c, <-c // 从通道c中接收

	fmt.Println(x, y, x+y)
}
