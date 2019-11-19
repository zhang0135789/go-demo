package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

/*
	库

	1. go开启了一个协成
	2. <- chan int   表示只能从 chan int 中拿东西  出去
*/

var startTime time.Time

func Init() {
	startTime = time.Now()
}

/**
把数据放入通道
*/
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		//发送方一定要close  如果不close 在range里面 就不知道什么时候结束
		close(out)
	}()
	return out
}

/**
对 通道内数据进行排序
*/
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read done:", time.Now().Sub(startTime))

		// 排序
		sort.Ints(a)
		fmt.Println("in mem sort done:", time.Now().Sub(startTime))

		//放入out
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

/**
合并 排序
*/
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("merge done:", time.Now().Sub(startTime))

	}()
	return out
}

/**
读取随机数
*/
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		count := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			//n 读到的字节数   err 是否有错误
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
				count++
				//fmt.Println(count)
				if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
					break
				}
			} else {
				break
			}
		}
		close(out)
		//fmt.Println("关闭通道")
	}()
	return out
}

/**
写入
*/
func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

/**
生成随机数
*/
func RadomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

/**
n个节点俩俩归并
*/
func MergeN(inputs ...<-chan int) <-chan int {
	size := len(inputs)
	if size == 1 {
		return inputs[0]
	}

	m := size / 2
	//merge inputs[0..m) adn inputs (m..end]
	return Merge(
		MergeN(inputs[:m]...), MergeN(inputs[m:]...))

}
