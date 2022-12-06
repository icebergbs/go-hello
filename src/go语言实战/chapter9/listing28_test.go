// 用来检测要将整数值转为字符串，使用那个函数会更好的基准
// 测试示例，先使用fmt.Sprintf函数，然后使用
// strconv.FormatInt函数，最后使用strconv.Itoa
package chapter9

import (
	"fmt"
	"strconv"
	"testing"
)

// BenchmarkSprintf 对 fmt.Sprintf函数
// 进行基准测试
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}

}

// BenchmarkFromate 对 strconv.FormatInf函数
// 进行基准测试
func BenchmarkFromate(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa 对 strconv.Itoa函数
// 进行基准测试
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}

/**
go test -v -run="none" -bench=.

*/
