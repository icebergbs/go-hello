package main

import "fmt"

//notifier 是一个定义了
// 通知类型行为的接口
type notifier6 interface {
	notify()
}

// user 在程序里定义一个用户类型
type user6 struct {
	name  string
	email string
}

// notify 实现了一个可以通过user类型值的指针
// 调用的方法
func (u *user6) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email,
		u.name,
		u.email)
}

// admin 代表一个拥有权限的管理员用户
type admin6 struct {
	user6 //嵌入类型
	level string
}

//main 是应用程序的入口
func main() {
	//创建一个admin 用户
	ad := admin6{
		user6: user6{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 给admin 用户发送一个通知
	// 用于实现接口的内部类型的方法，被提升到
	// 外部类型
	sendNotification6(&ad)
}

//sendNotification接受一个实现了notifier接口的值
// 并发送通知
func sendNotification6(n notifier6) {
	n.notify()
}
