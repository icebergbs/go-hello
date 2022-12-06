//这个示例实现了简单的网络服务
package main

import (
	handlers2 "go-hello/src/go语言实战/chapter91/handlers"
	"log"
	"net/http"
)

func main() {
	handlers2.Routes()
	log.Println("listener : Started : Listening on : 4000")
	http.ListenAndServe(":4000", nil)
}
