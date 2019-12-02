package main

import "fmt"

func main() {
	//demo5()

	//demo6()

	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := numbers4[4:6:8]
	length := (2)
	capacity := (4)
	fmt.Printf("%v, %v\n", length == len(slice5), capacity == cap(slice5))
	slice5 = slice5[:cap(slice5)]
	slice5 = append(slice5, 11, 12, 13)
	length = (7)
	fmt.Printf("%v\n", length == len(slice5))
	slice6 := []int{0, 0, 0}
	copy(slice5, slice6)
	e2 := (0)
	e3 := (8)
	e4 := (11)
	fmt.Printf("%v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])
}

func demo6() {
	number3 := [5]int{1, 2, 3, 4, 5}
	slice2 := number3[2:4:5]

	printArr(slice2)

	fmt.Printf("len  %d \n", len(slice2))
	fmt.Printf("cap  %d \n", cap(slice2))

}

func demo5() {
	number3 := [5]int{1, 2, 3, 4, 5}
	slice1 := number3[2:3]
	printArr(slice1)

	slice1 = slice1[0:3]
	printArr(slice1)

	slice1 = slice1[:cap(slice1)]
	printArr(slice1)

	fmt.Printf("  %d", cap(slice1))

}

func printArr(arrays []int) {
	for _, v := range arrays {
		fmt.Println(v)
	}
}
