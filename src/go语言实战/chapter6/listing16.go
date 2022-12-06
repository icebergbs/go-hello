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
	counter16 int

	// wg用来等待程序结束
	wg16 sync.WaitGroup

	// mutex 用来定义一段代码临界区
	mutex sync.Mutex
)

// main 是所有go程序的入口
func main() {

	// 计数加2， 表示要等待两个goroutine
	wg16.Add(2)

	//创建两个goroutine
	go incCounter16(1)
	go incCounter16(2)

	//等待goroutine结束
	wg16.Wait()
	fmt.Println("Final Counter:", counter16)

}

// incCounter 增加包里counter变量的值
func incCounter16(id int) {
	//在函数退出时调用Done 来通知main 函数工作已经完成
	defer wg16.Done()

	for count := 0; count < 2; count++ {
		//同一时刻只允许一个goroutine进入
		// 这个临界区
		mutex.Lock()
		{
			//捕获counter的值
			value := counter16

			//当前goroutine 从线程退出，并放回到队列
			runtime.Gosched()

			//增加本地value变量的值
			value++

			//将该值保存回counter
			counter16 = value
		}
		mutex.Unlock()
		//释放锁， 允许其他正在等待的goroutine
		// 进入临界区

	}
}
