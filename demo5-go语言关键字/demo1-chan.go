package main

import (
	"fmt"
	"time"
)

/*
	chan 通道
	对通道的读取和写入都可能进入阻塞状态。

	通道赋值
	1.不带缓冲的通道，在写入时，就会发生阻塞，直到通道中信息被读取后，才会结束阻塞。
	2.带缓冲的通道，每次向通道中写入一次信息，通道长度就会加1，每成功从通道读取一次信息，通道长度减1。如果通道长度等于通道缓冲长度时，向通道继续写入信息会使程序阻塞；如果通道长度小于通道缓冲长度，则向通道中写入信息不会造成阻塞。假如通道长度是5，那么在通道没有被读取的情况下，向通道中第6次写入信息时才会导致程序阻塞。

	读取通道
		通道为空
		1.通道没有关闭，程序会进入阻塞状态，等到通道有信息写入
		2.通道已经关闭，不会阻塞，返回通道中数据类型初始值（脏数据），如通道是chan int时，返回值是0，通道是chan string时，返回值是空。
		通道不为空
		1.通道没有关闭，从通道中读取一次信息，读取完成后，往下执行
		2.通道已被关闭，从通道中读取一次信信，读取完成后，往下执行
*/

func main() {
	//demo1()

	demo2()
}

func demo2() {

	// 构建一个通道
	ch := make(chan int) //构建一个同步用的通道
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine
		time.Sleep(time.Second)
		ch <- 0
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	// 等待匿名goroutine
	<-ch
	fmt.Println("all done")
}

func demo1() {
	// 定义一个不带缓冲的通道，通道中数据类型是int
	var c = make(chan int)

	// 开启一个携程，读取通道中的内容
	go func() {
		fmt.Println("写入信息是：", <-c)
	}()

	// 向通道中写入数据
	c <- 1
}
