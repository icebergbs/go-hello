package main

import "fmt"

func main() {

	//使用长度声明一个字符串切片
	// 创建一个字符串切片
	// 其长度和容量都是 5 个元素
	slice := make([]string, 5)
	fmt.Println("slice= ", slice)

	//使用长度和容量声明整型切片
	// 创建一个整型切片
	// 其长度为 3 个元素，容量为 5 个元素
	slice1 := make([]int, 3, 5)
	fmt.Println("slice1= ", slice1)

	//容量小于长度的切片会在编译时报错
	// 创建一个整型切片
	// 使其长度大于容量
	//slice2 := make([]int, 5, 3)
	//fmt.Println("slice1= ", slice2)

	//通过切片字面量来声明切片
	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println("slice3= ", slice3)
	// 创建一个整型切片
	// 其长度和容量都是 3 个元素
	slice4 := []int{10, 20, 30}
	fmt.Println("slice4= ", slice4)

	//使用索引声明切片
	// 创建字符串切片
	// 使用空字符串初始化第 100 个元素
	slice5 := []string{99: ""}
	fmt.Println("slice5= ", slice5)

	//声明数组和声明切片的不同
	// 创建有 3 个元素的整型数组
	array := [3]int{10, 20, 30}
	// 创建长度和容量都是 3 的整型切片
	slice6 := []int{10, 20, 30}
	fmt.Println("array= ", array)
	fmt.Println("slice6= ", slice6)

	//创建 nil 切片
	// 创建 nil 整型切片
	var slice7 []int
	fmt.Println("slice7= ", slice7)

	//声明空切片
	// 使用 make 创建空的整型切片
	slice8 := make([]int, 0)
	// 使用切片字面量创建空的整型切片
	slice9 := []int{}
	fmt.Println("slice8= ", slice8)
	fmt.Println("slice9= ", slice9)

	//使用切片字面量来声明切片
	// 创建一个整型切片
	// 其容量和长度都是 5 个元素
	slice11 := []int{10, 20, 30, 40, 50}
	// 改变索引为 1 的元素的值
	slice11[1] = 25
	fmt.Println("slice11= ", slice11)

	//使用切片创建切片
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice22 := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度为 2 个元素，容量为 4 个元素
	newSlice22 := slice22[1:3]
	fmt.Println("newSlice= ", newSlice22)

	//如何计算长度和容量
	//对底层数组容量是 k 的切片 slice[i:j]来说
	//长度: j - i
	//容量: k - i

	//修改切片内容可能导致的结果
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice33 := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度是 2 个元素，容量是 4 个元素
	newSlice33 := slice33[1:3]
	// 修改 newSlice 索引为 1 的元素
	// 同时也修改了原来的 slice 的索引为 2 的元素
	newSlice33[1] = 35
	fmt.Println("newSlice33= ", newSlice33)

	//表示索引越界的语言运行时错误
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	//slice44 := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度为 2 个元素，容量为 4 个元素
	//newSlice44 := slice44[1:3]
	// 修改 newSlice 索引为 3 的元素
	// 这个元素对于 newSlice 来说并不存在
	//newSlice44[3] = 45
	//fmt.Println("newSlice44= ", newSlice44)
	//切片有额外的容量是很好，但是如果不能把这些容量合并到切片的长度里，这些容量就没有
	//用处。好在可以用 Go 语言的内置函数 append 来做这种合并很容易。

	//使用 append 向切片增加元素
	// 其长度和容量都是 5 个元素
	slice55 := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度为 2 个元素，容量为 4 个元素
	newSlice55 := slice55[1:3]
	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为 60
	newSlice55 = append(newSlice55, 60)
	fmt.Println("slice55= ", slice55)
	fmt.Println("newSlice55= ", newSlice55)

	//使用 append 同时增加切片的长度和容量
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice66 := []int{10, 20, 30, 40}
	// 向切片追加一个新元素
	// 将新元素赋值为 50
	newSlice66 := append(slice66, 50)
	fmt.Printf("newSlice66= %s, len = %d\n", newSlice66, len(newSlice66))

	//使用切片字面量声明一个字符串切片
	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	slice77 := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	fmt.Println("slice77= ", slice77)

	//使用 3 个索引创建切片
	// 将第三个元素切片，并限制容量
	// 其长度为 1 个元素，容量为 2 个元素
	slice88 := slice77[2:3:4]
	fmt.Println("slice88= ", slice88)

	//如何计算长度和容量
	//对于 slice[i:j:k] 或 [2:3:4]
	//长度: j – i 或 3 - 2 = 1
	//容量: k – i 或 4 - 2 = 2

	//设置容量大于已有容量的语言运行时错误
	// 这个切片操作试图设置容量为 4
	// 这比可用的容量大
	//slice99 := slice77[2:3:6]
	//fmt.Println("slice99= ", slice99)

	//设置长度和容量一样的好处
	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	source111 := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 对第三个元素做切片，并限制容量
	// 其长度和容量都是 1 个元素
	slice111 := source111[2:3:3]
	// 向 slice 追加新字符串
	slice111 = append(slice111, "Kiwi")
	fmt.Println("source111 = ", source111)
	fmt.Println("slice111 = ", slice111)

	//将一个切片追加到另一个切片
	// 创建两个切片，并分别用两个整数进行初始化
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	// 将两个切片追加在一起，并显示结果
	fmt.Printf("s1 append s2 = %v\n", append(s1, s2...))

	//使用 for range 迭代切片
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice222 := []int{10, 20, 30, 40}
	// 迭代每一个元素，并显示其值
	for index, value := range slice222 {
		fmt.Printf("for range Index: %d Value: %d\n", index, value)
	}

	//range 提供了每个元素的副本
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice333 := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slice333 {
		fmt.Printf("for range Value: %d Value-Addr: %X Element-Addr: %X\n",
			value, &value, &slice333[index])
	}

	//使用空白标识符（下划线）来忽略索引值
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice444 := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示其值
	for _, value := range slice444 {
		fmt.Printf("Value: %d\n", value)
	}

	//使用传统的 for 循环对切片进行迭代
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice555 := []int{10, 20, 30, 40}
	// 从第三个元素开始迭代每个元素
	for index := 2; index < len(slice555); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice555[index])
	}

	//声明多维切片
	// 创建一个整型切片的切片
	slice1111 := [][]int{{10}, {100, 200}}
	fmt.Printf("slice1111: %d\n", slice1111)

	//组合切片的切片
	// 创建一个整型切片的切片
	slice2222 := [][]int{{10}, {100, 200}}
	// 为第一个切片追加值为 20 的元素
	slice2222[0] = append(slice2222[0], 20)
	fmt.Printf("slice2222: %d\n", slice2222)

	//在函数间传递切片
	// 分配包含 100 万个整型值的切片
	slice3333 := make([]int, 100)
	// 将 slice 传递到函数 foo
	slice3333 = foo3(slice3333)
	fmt.Printf("slice3333: %d\n", slice3333)

	//4.3 映射的内部实现和基础功能

}

// 函数 foo 接收一个整型切片，并返回这个切片
func foo3(slice []int) []int {
	return slice
}
