package main

import (
	"fmt"
	"go-hello/src/go语言学习指南/c9模块包和导入/package_example/formatter"
	"go-hello/src/go语言学习指南/c9模块包和导入/package_example/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
