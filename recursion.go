package main

import (
	"fmt"
	"time"
)

// 阶乘
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// 斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int = 15
	fmt.Println("i的uint64是：", uint64(i))
	fmt.Println("开始时间：", time.Now().UnixNano())
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
	fmt.Println("结束时间：", time.Now().UnixNano())

	fmt.Println("开始时间：", time.Now().Unix())
	var j int
	for j = 0; j < 40; j++ {
		fmt.Printf("%d\t", fibonacci(j))
	}

	fmt.Println("\n结束时间：", time.Now().Unix())

}
