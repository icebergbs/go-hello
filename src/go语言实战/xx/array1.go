package main

import "fmt"

func main() {
	/**
	4.1.2 声明和初始化
	*/

	// 声明一个包含 5 个元素的整型数组
	var array [5]int
	fmt.Println("arr = %s", array)

	// 声明一个包含 5 个元素的整型数组
	// 用具体值初始化每个元素
	array1 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("arr1 = %s", array1)

	// 用具体值初始化每个元素
	// 容量由初始化值的数量决定
	array2 := [...]int{10, 20, 30, 40, 50}
	fmt.Println("arr2 = %s", array2)

	// 用具体值初始化索引为 1 和 2 的元素
	// 其余元素保持零值
	array3 := [5]int{1: 10, 2: 20}
	fmt.Println("arr3 = %s", array3)

	//访问数组元素
	// 声明一个包含 5 个元素的整型数组
	// 用具体值初始为每个元素
	array4 := [5]int{10, 20, 30, 40, 50}
	// 修改索引为 2 的元素的值
	array4[2] = 35
	fmt.Println("arr4 = %s", array4)

	//访问指针数组的元素
	// 声明包含 5 个元素的指向整数的数组
	// 用整型指针初始化索引为 0 和 1 的数组元素
	array5 := [5]*int{0: new(int), 1: new(int)}
	// 为索引为 0 和 1 的元素赋值
	*array5[0] = 10
	*array5[1] = 20
	fmt.Println("arr5[0] = %s", *array5[0])
	fmt.Println("arr5[1] = %s", *array5[1])
	fmt.Println("arr5 = %s", array5)

	//把同样类型的一个数组赋值给另外一个数组
	// 声明第一个包含 5 个元素的字符串数组
	var array6 [5]string
	// 声明第二个包含 5 个元素的字符串数组
	// 用颜色初始化数组
	array7 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// 把 array2 的值复制到 array1
	array6 = array7
	fmt.Println("arr6 = arr7 :", array6 == array7)

	//编译器会阻止类型不同的数组互相赋值
	// 声明第一个包含 4 个元素的字符串数组
	//var array8 [4]string
	//// 声明第二个包含 5 个元素的字符串数组
	//// 使用颜色初始化数组
	//array9 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	//// 将 array2 复制给 array1
	//array8 = array9

	//把一个指针数组赋值给另一个
	// 声明第一个包含 3 个元素的指向字符串的指针数组
	var array11 [3]*string
	// 声明第二个包含 3 个元素的指向字符串的指针数组
	// 使用字符串指针初始化这个数组
	array22 := [3]*string{new(string), new(string), new(string)}
	// 使用颜色为每个元素赋值
	*array22[0] = "Red"
	*array22[1] = "Blue"
	*array22[2] = "Green"
	// 将 array2 复制给 array1
	array11 = array22
	fmt.Println("arr11= ", array11)
	fmt.Println("arr22= ", array22)

	//声明二维数组
	// 声明一个二维整型数组，两个维度分别存储 4 个元素和 2 个元素
	var array33 [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array33 = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println("arr33= ", array33)
	// 声明并初始化外层数组中索引为 1 个和 3 的元素
	array33 = [4][2]int{1: {20, 21}, 3: {40, 41}}
	fmt.Println("arr33= ", array33)
	// 声明并初始化外层数组和内层数组的单个元素
	array33 = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println("arr33= ", array33)

	//访问二维数组的元素
	// 声明一个 2×2 的二维整型数组
	var array44 [2][2]int
	// 设置每个元素的整型值
	array44[0][0] = 10
	array44[0][1] = 20
	array44[1][0] = 30
	array44[1][1] = 40
	fmt.Println("arr44= ", array44)

	//同样类型的多维数组赋值
	// 声明两个不同的二维整型数组
	var array55 [2][2]int
	var array66 [2][2]int
	// 为每个元素赋值
	array66[0][0] = 10
	array66[0][1] = 20
	array66[1][0] = 30
	array66[1][1] = 40
	// 将 array2 的值复制给 array1
	array55 = array66
	fmt.Println("arr55= ", array55)
	fmt.Println("arr66= ", array66)

	//使用索引为多维数组赋值
	// 将 array1 的索引为 1 的维度复制到一个同类型的新数组里
	var array77 [2]int = array66[1]
	fmt.Println("arr77= ", array77)
	// 将外层数组的索引为 1、内层数组的索引为 0 的整型值复制到新的整型变量里
	var value int = array66[1][0]
	fmt.Println("value= ", value)

	//使用值传递，在函数间传递大数组
	// 声明一个需要 8 MB 的数组
	var array88 [8]int
	// 将数组传递给函数 foo
	foo(array88)

	//使用指针在函数间传递大数组
	// 分配一个需要 8 MB 的数组
	var array99 [8]int
	// 将数组的地址传递给函数 foo
	foo1(&array99)

}

// 函数 foo 接受一个 100 万个整型值的数组
func foo(array [8]int) {
	fmt.Println("array= ", array)
}

// 函数 foo 接受一个指向 100 万个整型值的数组的指针
func foo1(array *[8]int) {
	fmt.Println("*array= ", array)
}
