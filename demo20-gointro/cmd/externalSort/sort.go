package main

import (
	"../pipeline"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
外部排序 单机版 pipeline
ReaderSource节点  - 支持分块
MergeN - 搭建归并节点组
pipeline的搭建及运行  - CPU及县城数量的观测
*/
func main() {
	//demo1()

	demo2()

}

func demo2() {
	infilename := "large.in"
	//outfilename := "large.out"
	p := createNetworkPipeline(infilename, 800000000, 8)

	writeToFile(p, "large2.out")

	printFile("large2.out")
}

func demo1() {
	infilename := "large.in"
	outfilename := "large.out"
	p := createPipeline(infilename, 800000000, 8)

	writeToFile(p, outfilename)

	printFile(outfilename)
}

/*
	读取文件
*/
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	source := pipeline.ReaderSource(bufio.NewReader(file), -1)

	count := 0
	for v := range source {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}
}

/**
写入文件
*/
func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	defer writer.Flush()

}

/**
生成pipel
filename  文件名
chunkCount 文件切分块数
*/
func createPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipeline.Init()

	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}

func createNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipeline.Init()

	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)

		pipeline.NetworkSink(addr, pipeline.InMemSort(source))

		sortAddr = append(sortAddr, addr)

	}

	sortResults := []<-chan int{}

	//链接 server
	// 取出个节点的结果  进行合并
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}

	return pipeline.MergeN(sortResults...)
}
