package main

import (
	"fmt"
)

// notifier 是一个定义了
// 通知类行为的接口
type notifier interface {
	notify()
}

// user 在程序里定义一个用户类型
type user2 struct {
	name  string
	email string
}

// notify 是使用指针接受者实现的方法
func (u *user2) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main 是应用程序的入口
func main() {
	//创建一个user类型的值， 并发送通知
	u := user2{"Bill", "bill@eamil.com"}

	//sendNotification(u)
	//Cannot use 'u' (type user) as type notifier Type does not implement 'notifier' as 'notify' method has a pointer receiver

	sendNotification(&u)

}

// sendNotification接受一个实现了notifier接口的值
// 并发送通知
func sendNotification(n notifier) {
	n.notify()
}
