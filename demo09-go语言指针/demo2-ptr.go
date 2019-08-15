package main

import "fmt"

/*
	ptr 为整型指针数组。因此每个元素都指向了一个值。
*/
const MAX int = 3

func main() {
	a := []int{10, 100, 200}
	//var b = [3]int{1,2,3}
	fmt.Println(a)
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
