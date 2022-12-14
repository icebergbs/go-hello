//这个示例程序使用接口展示多台行为
package main

import "fmt"

// notifier 是一个定义了
// 通知类行为的接口
type notifier3 interface {
	notify()
}

// user 在程序里定义一个用户类型
type user4 struct {
	name  string
	email string
}

//notify 使用指针接受者实现了notifier接口
func (u *user4) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin 定义了程序里的管理员
type admin1 struct {
	name  string
	email string
}

// notify 使用指针接受者实现了notifier接口
func (a *admin1) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

// main 是应用程序的入口
func main() {
	//创建一个user值并传给sendNotification
	bill := user4{"Bill", "bill@email.com"}
	sendNotification3(&bill)

	// 创建一个admin 值并传给sendNotification
	lisa := admin1{"Lisa", "lisa@email.com"}
	sendNotification3(&lisa)
}

//sendNotification接受一个实现了notifier接口的值
// 并发送通知
func sendNotification3(n notifier3) {
	n.notify()
}
