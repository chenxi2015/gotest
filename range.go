package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	//var sum int = 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for index, value := range kvs {
		fmt.Printf("index: %s, value: %s\n", index, value)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for int, c := range "go" {
		fmt.Printf("int: %d, c: %d\n", int, c)
	}

	for int, c := range "ABC" {
		fmt.Printf("index: %d, c: %d\n", int, c)
	}
}
