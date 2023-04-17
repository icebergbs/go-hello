package main

import (
	"fmt"
	"go-hello/src/go语言学习指南/c9模块包和导入/circular_dependency_example/person"
)

func main() {
	bob := person.Person{PetName: "Fluffy"}
	fmt.Println(bob.Pet())
}
