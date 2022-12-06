//这个示例程序展示如何在程序里造成竞争状态
// 实际上不希望出现这种情况
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	//counter是所有goroutine都要增加其值的变量
	counter int

	// wg用来等待程序结束
	wg9 sync.WaitGroup
)

// main 是所有go程序的入口
func main() {

	// 计数加2， 表示要等待两个goroutine
	wg9.Add(2)

	//创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	//等待goroutine结束
	wg9.Wait()
	fmt.Println("Final Counter:", counter)

}

// incCounter 增加包里counter变量的值
func incCounter(id int) {
	//在函数退出时调用Done 来通知main 函数工作已经完成
	defer wg9.Done()

	for count := 0; count < 2; count++ {
		//捕获counter的值
		value := counter

		//当前goroutine 从线程退出，并放回到队列
		runtime.Gosched()

		//增加本地value变量的值
		value++

		//将该值保存回counter
		counter = value
	}
}
