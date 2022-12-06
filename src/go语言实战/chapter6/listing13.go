//这个示例程序展示如何使用atomic包来提供
// 对数值类型的安全访问
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	//counter是所有goroutine都要增加其值的变量
	counter13 int64

	// wg用来等待程序结束
	wg13 sync.WaitGroup
)

// main 是所有go程序的入口
func main() {

	// 计数加2， 表示要等待两个goroutine
	wg13.Add(2)

	//创建两个goroutine
	go incCounter13(1)
	go incCounter13(2)

	//等待goroutine结束
	wg13.Wait()
	fmt.Println("Final Counter:", counter13)

}

// incCounter 增加包里counter变量的值
func incCounter13(id int) {
	//在函数退出时调用Done 来通知main 函数工作已经完成
	defer wg13.Done()

	for count := 0; count < 2; count++ {
		//安全地对counter加1
		atomic.AddInt64(&counter13, 1)

		//当前goroutine 从线程退出，并放回到队列
		runtime.Gosched()
	}
}
