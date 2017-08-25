package main

import (
	"fmt"
)
// s[lo:hi] 表示从 lo 到 hi-1 的 slice 元素，含前端，不包含后端。因此

func main()  {
	s := []int {2,3,5,7,9,10}

	fmt.Println(s)
	fmt.Println("s[1:4] = ", s[1:4])

	fmt.Println("s[:3] = ", s[:3])

	fmt.Println("s[4:] = ", s[4:])
	fmt.Println("s[1:1] = ", s[1:1])
}