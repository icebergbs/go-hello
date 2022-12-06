// 这个示例程序展示如何使用io.Reader和io.Writer接口
// 写一个简单版本的curl
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	arr1 := [2]string{"http://localhost:7009/interface/void/getLabelInfo", "D:\\data\\logs"}
	// 这里的r 是一个响应， r.Body是io.Reader
	r, err := http.Get(arr1[0])
	if err != nil {
		log.Fatalln(err)
	}

	// 创建文件来保存响应内容
	file, err := os.Create(arr1[1])
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	// 使用MultiWriter, 这样就可以同时向文件和标准输出设备
	// 进行写操作
	dest := io.MultiWriter(os.Stdout, file)

	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}