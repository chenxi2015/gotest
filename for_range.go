package main

import (
	"fmt"
	"time"
)

// for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。

// 当使用 for 循环遍历一个 slice 时，每次迭代 range 将返回两个值。 第一个是当前下标（序号），第二个是该下标所对应元素的一个拷贝。

// 可以通过赋值给 _ 来忽略序号和值。

// 如果只需要索引值，去掉 “ , value ” 的部分即可。

var pow = []int{1,2,4,8,16,64,128}
func main()  {
	fmt.Println("starttme:", time.Now().UnixNano())
	starttime := time.Now().UnixNano();
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10)

	for i := range pow {
		// fmt.Println("i == ", i)
		pow[i] = 1 << uint(i) // 左移 1 01 左移 2位  0100
	}
	fmt.Println("pow", pow)
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	fmt.Println("endtme:", time.Now().UnixNano())
	fmt.Println("usetime:", (time.Now().UnixNano() - starttime) / 1000 * 1000)
	
}