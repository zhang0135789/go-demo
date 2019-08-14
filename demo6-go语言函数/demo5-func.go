package main

import "fmt"

/*
	Go 语言函数闭包
	Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明
*/

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence() //将getSequence()方法返回 的函数 赋值给nextNumber

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
