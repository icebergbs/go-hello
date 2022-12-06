// 这个示例程序展示如何解码JSON字符串
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//
type Contact1 struct {
	Name string `json:"name""`
	Title string `json:"title""`
	Contact1 struct{
		Home string `json:"home""`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// JSON contains a sample string to unmarshal.
var JSON1 = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	//将JSON字符串反序列化到map变量
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON1), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Name", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])
}