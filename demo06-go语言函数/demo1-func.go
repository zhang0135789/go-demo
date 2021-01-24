package main

import "fmt"

/*
	函数定义
	func function_name( [parameter list] ) [return_types] {
   		函数体
	}
*/
func main() {
	/* 定义局部变量 */
	//demo1()

	/*函数返回多个值*/
	demo2()
}

/*
 	a=1
	b=2

=>
  	a=2
	b=1

c
*/
func demo2() {
	a, b := swap("google", "runoob")
	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}

/*

 */
func demo1() {
	var a int = 100
	var b int = 200
	var ret int
	/* 调用函数并返回最大值 */
	ret = max(a, b)
	fmt.Printf("最大值是 : %d\n", ret)
}

func max(num1 int, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
