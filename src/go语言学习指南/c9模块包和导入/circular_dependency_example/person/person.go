package person

import "go-hello/src/go语言学习指南/c9模块包和导入/circular_dependency_example/pet"

type Person struct {
	Name    string
	Age     int
	PetName string
}

var pets = map[string]pet.Pet{
	"Fluffy": {"Fluffy", "Cat", "Bob"},
	"Rex":    {"Rex", "Dog", "Julia"},
}

func (p Person) Pet() pet.Pet {
	return pets[p.PetName]
}
