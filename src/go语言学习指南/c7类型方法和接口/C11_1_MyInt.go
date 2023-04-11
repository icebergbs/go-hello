package main

import (
	"fmt"
	"io"
)

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	//i2 := i.(MyInt)
	//fmt.Println(i2 + 1) // 21

	//i2 := i.(string) // panic:
	//fmt.Println(i2) // 21

	//i2 := i.(int) // panic:
	//fmt.Println(i2) // 21

	//ok模式避免程序崩溃
	i2, ok := i.(int) // panic:
	if !ok {
		fmt.Errorf("unexpected type for %v", i)
		return
	}
	fmt.Println(i2 + 1)

}

//当一个接口是多种可能的类型之一时，使用type-switch来代替：
func doThings(i interface{}) {
	switch j := i.(type) {
	case nil:
		// i is nil, type of j is interface{}
	case int:
		// j is of type int
		fmt.Println(j)
	case MyInt:
		// j is of type MyInt
	case io.Reader:
		// j is of type io.Reader
	case string:
		// j is a string
	case bool, rune:
		// i is either a bool or rune, so j is of type interface{}
	default:
		// no idea what i is, so j is of type interface{}

	}
}
