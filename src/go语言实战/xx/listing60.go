package main

import "fmt"

//notifier 是一个定义了
// 通知类型行为的接口
type notifier7 interface {
	notify()
}

// user 在程序里定义一个用户类型
type user7 struct {
	name  string
	email string
}

// notify 实现了一个可以通过user类型值的指针
// 调用的方法
func (u *user7) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email,
	)
}

// admin 代表一个拥有权限的管理员用户
type admin7 struct {
	user7 //嵌入类型
	level string
}

// 通过admin 类型值的指针
// 调用的方法
func (a *admin7) notify() {
	fmt.Printf("sending admin level to %s\n",
		a.level)
}

//main 是应用程序的入口
func main() {
	//创建一个admin 用户
	ad := admin7{
		user7: user7{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 给admin 用户发送一个通知
	// 用于实现接口的内部类型的方法，被提升到
	// 外部类型
	sendNotification7(&ad)

	//我们可以直接访问内部类型的方法
	ad.user7.notify()

	//内部类型的方法没有被提升
	ad.notify()
}

//sendNotification接受一个实现了notifier接口的值
// 并发送通知
func sendNotification7(n notifier7) {
	n.notify()
}
