package main

import "fmt"

func main() {
	var x = 10
	var i *int = new(int)
	*i = x
	fmt.Println(*i)
}
