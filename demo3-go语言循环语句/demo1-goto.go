package main

import "fmt"

func main() {

	//demo1_goto()

	demo2_for()
}

/*
for key, value := range oldMap {
    newMap[key] = value
}
*/
func demo2_for() {

	var b int = 15
	var a int = 1
	numbers := [6]int{1, 2, 3, 5}
	/* for 循环 */
	for a := 0; a < 10; a++ {
		a++
		fmt.Printf("a = 的值为: %d\n", a)
	}
	fmt.Printf("%d\n", a)
	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}
	fmt.Printf("%d\n", a)
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

/*
   goto label;
   ..
   .
   label: statement;
*/
func demo1_goto() {
	/* 定义局部变量 */
	var a int = 10
	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
