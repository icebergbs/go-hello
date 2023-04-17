package sibling

import "go-hello/src/go语言学习指南/c9模块包和导入/internal_package_example/foo/internal"

func AlsoUseDoubler(i int) int {
	return internal.Double(i)
}
