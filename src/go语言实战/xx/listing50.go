package main

import "fmt"

// user 在程序里定义一个用户类型
type user5 struct {
	name  string
	email string
}

// notify 实现了一个可以通过user类型值的指针
// 调用的方法
func (u *user5) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)

}

// admin 代表一个拥有权限的管理员用户
type admin5 struct {
	user5 //嵌入类型
	level string
}

//main 是应用程序的入口
func main() {
	//创建一个admin 用户
	ad := admin5{
		user5: user5{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 我们可以直接访问内部类型的方法
	ad.user5.notify()

	//内部类型的方法也被提升到外部类型
	ad.notify()
}
