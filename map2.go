package main

import (
	"fmt"
	"time"
)

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "印度"
	countryCapitalMap["China"] = "中国"

	fmt.Println("开始时间：", time.Now().UnixNano())

	for country := range countryCapitalMap {
		fmt.Printf("country: %s, name: %s\n", country, countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["America"]
	if ok {
		fmt.Println("美国的首都是： ", capital)
	} else {
		fmt.Println("不存在美国首都")
	}
	fmt.Println("结束时间：", time.Now().UnixNano())

}
