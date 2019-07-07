package main

import "fmt"

func main() {
	countryCapitalMap := map[string]string{"France": "巴黎", "Italy": "新德里", "China": "北京"}
	fmt.Println("原始地图")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是：", countryCapitalMap[country])
	}

	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后的地图")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是：", countryCapitalMap[country])
	}
}
