package main

import "fmt"

/*
	全局变量
	Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑
*/

/* 声明全局变量 */
var g int

func main() {

	test1()

	test2()
}

func test2() {
	/* 声明局部变量 */
	var g int = 10

	fmt.Printf("结果： g = %d\n", g)
}

func test1() {
	/* 声明局部变量 */
	var a, b int
	/* 初始化参数 */
	a = 10
	b = 20
	g = a + b
	fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)
}
