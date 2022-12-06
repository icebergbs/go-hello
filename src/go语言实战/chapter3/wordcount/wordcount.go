package main

import (
	"fmt"
	"go-hello/src/go语言实战/chapter3/words"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Therr was an error opening the file:", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
