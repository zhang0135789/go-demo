package main

import "fmt"

func main() {
	/*
		变量声明
	*/

	/*第一种，指定变量类型，如果没有初始化，则变量默认为零值。*/
	var a = "RUNOOB"
	fmt.Println(a)

	//int 默认为 0
	var b int
	fmt.Println(b)

	//bool 默认为false
	var c bool
	fmt.Println(c)
}
