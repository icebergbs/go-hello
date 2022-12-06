package main

import "fmt"

func main() {
	//使用 make 创建通道
	// 无缓冲的整型通道
	//unbuffered := make(chan int)
	// 有缓冲的字符串通道
	buffered := make(chan string, 10)

	//向通道发送值
	// 通过通道发送一个字符串
	buffered <- "Gopher"
	//unbuffered <- 1

	//从通道里接收值
	// 从通道接收一个字符串
	value := <-buffered
	fmt.Println("value = ", value)

}
