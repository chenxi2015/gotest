package main

import "fmt"

func main() {

	var numbers []int
	printSlice5(numbers)

	// append
	numbers = append(numbers, 0)
	printSlice5(numbers)

	numbers = append(numbers, 1)
	printSlice5(numbers)

	numbers = append(numbers, 2, 3, 4, 5)
	printSlice5(numbers)

	//var number1 []int = make([]int, len(numbers), cap(numbers) * 2)
	number1 := make([]int, len(numbers), cap(numbers)*2)
	printSlice5(number1)

	copy(number1, numbers)
	printSlice5(number1)

}
func printSlice5(x []int) {
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(x), cap(x), x)
}
