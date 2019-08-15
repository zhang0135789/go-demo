package main

import "fmt"

/*
	Go 语言切片(Slice)
	Go 语言切片是对数组的抽象。

	定义切片
	var identifier []type
	用make()函数来创建切片
	var slice1 []type = make([]type, len)
	也可以简写为
	slice1 := make([]type, len)
	也可以指定容量，其中capacity为可选参数。
	make([]T, length, capacity)

	切片初始化
	s :=[] int {1,2,3 }
	直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	s := arr[:]
	初始化切片s,是数组arr的引用
	s := arr[startIndex:endIndex]
	将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
	s := arr[startIndex:]
	默认 endIndex 时将表示一直到arr的最后一个元素
	s := arr[:endIndex]
	默认 startIndex 时将表示从arr的第一个元素开始
	s1 := s[startIndex:endIndex]
	通过切片s初始化切片s1
	s :=make([]int,len,cap)
	通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片




*/
func main() {
	//demo1()

	//demo2()

	//demo3()

	demo4()
}

/*
append() 和 copy() 函数
*/
func demo4() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	printSlice(numbers1)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

/*
切片截取
*/
func demo3() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}

/*
空(nil)切片
*/
func demo2() {
	var numbers []int

	printSlice(numbers)

	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

/*
len() 和 cap() 函数
*/
func demo1() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
