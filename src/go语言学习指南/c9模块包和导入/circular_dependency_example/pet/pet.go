package pet

import "go-hello/src/go语言学习指南/c9模块包和导入/circular_dependency_example/person"

type Pet struct {
	Name      string
	Type      string
	OwnerName string
}

var owners = map[string]person.Person{
	"Bob":   {"Bob", 30, "Fluffy"},
	"Julia": {"Julia", 40, "Rex"},
}

func (p Pet) Owner() person.Person {
	return owners[p.OwnerName]
}
