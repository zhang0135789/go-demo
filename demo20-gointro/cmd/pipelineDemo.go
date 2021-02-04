package main

import (
	"./pipeline"
	"bufio"
	"fmt"
	"os"
)

/**
排序案例
*/
func main() {

	//demo1()

	//demo2()

	//demo3()

	demo4()

}

/**
读取文件进行排序
*/
func demo4() {
	const filename = "small.in"
	const n = 64

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	//类似与Java的  finally
	defer file.Close()

	p := pipeline.RadomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//reader := bufio.NewReader(file)
	p = pipeline.ReaderSource(file, -1)
	//
	//fmt.Println("结束2")
	//
	//count := 0
	for v := range p {
		fmt.Println(v)
		//count ++
		//fmt.Println(count)

		//if count >= n {
		//	break
		//}
	}

	fmt.Println("结束")
}

/**
合并排序
*/
func demo3() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(
				2, 1, 5, 7, 65, 3)),
		pipeline.InMemSort(
			pipeline.ArraySource(
				7, 4, 0, 3, 2, 13, 8)))

	print2(p)
}

func demo2() {
	p := pipeline.InMemSort(pipeline.ArraySource(2, 1, 5, 7, 65, 3))

	print2(p)
}

func demo1() {
	p := pipeline.ArraySource(2, 1, 5, 7, 65, 3)

	//通道的两种遍历方法
	//写法一
	print1(p)

	//写法二
	//print2(p)
}

func print1(p <-chan int) {
	//写法一
	for {
		if num, ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}
	}
}

func print2(p <-chan int) {
	for v := range p {
		fmt.Println(v)
	}

}
