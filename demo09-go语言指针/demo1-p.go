package main

import "fmt"

/*
	Go 语言指针
	1.Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
	2.一个指针变量指向了一个值的内存地址。



*/
func main() {
	demo1()

	//demo2()
}

/*
Go 空指针
*/
func demo2() {
	var ptr *int
	fmt.Printf("ptr 的值为 : %d\n", ptr)
	fmt.Printf("ptr 的值为 : %v\n", ptr != nil)
	fmt.Printf("ptr 的值为 : %v\n", ptr == nil)
}

/*
使用指针
*/
func demo1() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}
