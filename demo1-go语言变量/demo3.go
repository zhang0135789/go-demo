package main

import "fmt"

func main() {
	/*
		数值类型（包括complex64/128）为 0

		布尔类型为 false

		字符串为 ""（空字符串）

		第一种，指定变量类型，如果没有初始化，则变量默认为零值。
		以下几种类型为 nil：

		var a *int
		var a []int
		var a map[string] int
		var a chan int
		var a func(string) int
		var a error // error 是接口
	*/
	demo1()

	/*
		第二种，根据值自行判定变量类型。
	*/
	demo2()

	/*
		第三种，省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式： v_name := value
	*/
	demo3()

}

func demo3() {
	f := "Runoob" // var f string = "Runoob"
	fmt.Println(f)
}

func demo2() {
	var d = true
	fmt.Println(d)
}

func demo1() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println("%v %v %v %q\n", i, f, b, s)
}
