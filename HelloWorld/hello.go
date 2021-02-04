package main

import (
	"fmt"
	util "hello.go/hello"
)

func main() {

	util.PrintHello()

	fmt.Println("Hello , World!")

	char1 := '赞'
	fmt.Printf("字符 %c 的Unicode代码点是 %U  \n", char1, char1)

	var numbers3 = [5]int{1, 2, 3, 4, 5}
	var slice1 = numbers3[1:4]
	for _, v := range slice1 {
		fmt.Printf("%d \n", v)
	}

}

//執行命令  go run hello.go
