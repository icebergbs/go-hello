package main

import "fmt"

func main() {

	//使用 make 声明映射
	// 创建一个映射，键的类型是 string，值的类型是 int
	dict := make(map[string]int)
	fmt.Println("dict = ", dict)
	// 创建一个映射，键和值的类型都是 string
	// 使用两个键值对初始化映射
	dict1 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println("dict1 = ", dict1)

	//使用映射字面量声明空映射
	// 创建一个映射，使用字符串切片作为映射的键
	//dict2 := map[[]string]int{}

	//声明一个存储字符串切片的映射
	// 创建一个映射，使用字符串切片作为值
	dict3 := map[int][]string{}
	fmt.Println("dict3 = ", dict3)

	//为映射赋值
	// 创建一个空映射，用来存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{}
	// 将 Red 的代码加入到映射
	colors["Red"] = "#da1337"
	fmt.Println("colors = ", colors)

	//对 nil 映射赋值时的语言运行时错误
	// 通过声明映射创建一个 nil 映射
	//var colors1 map[string]string
	// 将 Red 的代码加入到映射
	//colors1["Red"] = "#da1337"

	//从映射获取值并判断键是否存在
	// 获取键 Blue 对应的值
	value, exists := colors["Blue"]
	// 这个键存在吗？
	if exists {
		fmt.Println(value)
	}

	//从映射获取值，并通过该值判断键是否存在
	// 获取键 Blue 对应的值
	value1 := colors["Blue"]
	// 这个键存在吗？
	if value1 != "" {
		fmt.Println(value1)
	}

	//使用 range 迭代映射
	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors2 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range colors2 {
		fmt.Printf("range Key: %s Value: %s\n", key, value)
	}

	//从映射中删除一项
	// 删除键为 Coral 的键值对
	delete(colors2, "Coral")
	// 显示映射里的所有颜色
	for key, value := range colors2 {
		fmt.Printf("delete Key: %s Value: %s\n", key, value)
	}

	//在函数间传递映射
	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors3 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range colors3 {
		fmt.Printf("function before Key: %s Value: %s\n", key, value)
	}
	// 调用函数来移除指定的键
	removeColor(colors3, "Coral")
	// 显示映射里的所有颜色
	for key, value := range colors3 {
		fmt.Printf("function after Key: %s Value: %s\n", key, value)
	}

}

// removeColor 将指定映射里的键删除
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
