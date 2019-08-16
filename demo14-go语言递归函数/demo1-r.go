package main

import "fmt"

/*
	Go 语言递归函数
*/
func main() {
	demo1()

	demo2()
}

func demo2() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

func demo1() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

/*
	阶乘
*/
func Factorial(n uint64) (result uint64) {
	if n > 1 {
		result = n * Factorial(n-1)
	} else {
		result = 1
	}
	return result
}

/*
	斐波那契数列
*/
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
