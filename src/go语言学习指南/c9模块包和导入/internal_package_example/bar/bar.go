package bar

import "go-hello/src/go语言学习指南/c9模块包和导入/internal_package_example/foo/internal"

func FailUseDoubler(i int) int {
	return internal.Doubler(i)
}
