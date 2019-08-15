package main

import "fmt"

/*
	Go 语言Map(集合)
	Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

	定义 Map
	// 声明变量，默认 map 是 nil /
	var map_variable map[key_data_type]value_data_type

	// 使用 make 函数 /
	map_variable := make(map[key_data_type]value_data_type)

	如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
*/
func main() {

	//demo1()

	demo2()
}

/*
	delete() 函数
	delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。实例如下：
*/
func demo2() {
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}

/*

 */
func demo1() {
	/*创建集合 */
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"]
	/*如果确定是真实的,则存在,否则不存在 */
	fmt.Println(capital)
	fmt.Println(ok)
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}
