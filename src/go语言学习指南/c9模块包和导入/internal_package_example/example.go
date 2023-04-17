package internal_package_example

import "go-hello/src/go语言学习指南/c9模块包和导入/internal_package_example/foo/internal"

func FailUseDoubler(a int) int {
	return internal.Doubler(a)
}
