package main

import (
	"fmt"
)
// map 映射键到值。

// map 在使用之前必须用 make 来创建；值为 nil 的 map 是空的，并且不能对其赋值。

type Vertex struct {

	Lat, Lang float64
}

var m map[string]Vertex

func main()  {
	m = make(map[string]Vertex)
	m["chenxi"] = Vertex{
		40.12312, 117.12312,
	}	
	fmt.Println(m["chenxi"])
	fmt.Println(m)
}