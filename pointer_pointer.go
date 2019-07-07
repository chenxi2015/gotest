package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	ptr = &a

	pptr = &ptr

	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("变量地址 ptr = %d\n", ptr)
	fmt.Printf("指向变量地址的变量地址 pptr = %d\n", pptr)
	fmt.Printf("指针变量 ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 pptr = %d\n", *pptr)
	fmt.Printf("指向指针的指针变量 pptr = %d\n", **pptr)
}
