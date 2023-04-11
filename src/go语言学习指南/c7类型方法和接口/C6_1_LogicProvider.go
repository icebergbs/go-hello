package main

import "fmt"

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// business logic
	fmt.Println("process data")
	return "process data"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	//get data from somewhere
	c.L.Process("data")
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
