package main

import "fmt"

//声明一个结构类型
// user 在程序里定义一个用户类型
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

//使用其他结构类型声明字段
// admin 需要一个 user 类型作为管理者，并附加权限
type admin struct {
	person user
	level  string
}

type Duration int64

func main() {

	//使用结构类型声明变量，并初始化为其零值
	// 声明 user 类型的变量
	var bill user
	fmt.Println("struct bill = ", bill)

	//使用结构字面量来声明一个结构类型的变量
	// 声明 user 类型的变量，并初始化所有字段
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}
	fmt.Println("struct lisa = ", lisa)

	//使用结构字面量创建结构类型的值
	//user{
	//	 name: "Lisa",
	//	 email: "lisa@email.com",
	//	 ext: 123,
	//	 privileged: true,
	//}
	//不使用字段名，创建结构类型的值
	// 声明 user 类型的变量
	lisa1 := user{"Lisa", "lisa@email.com", 123, true}
	fmt.Println("struct lisa1 = ", lisa1)

	//使用结构字面量来创建字段的值
	// 声明 admin 类型的变量
	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}
	fmt.Println("struct fred = ", fred)

	//给不同类型的变量赋值会产生编译错误
	//var dur Duration
	//dur = int64(1000)

}
