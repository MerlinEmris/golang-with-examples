package main

import "fmt"

type Employee interface {
	GetName() string
}

type ITEngineer struct {
	Name string
}

func (e *ITEngineer) GetName() string {
	return "E name: " + e.Name
}

func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}

func main() {
	engineer := &ITEngineer{Name: "Bashim"}
	PrintDetails(engineer)
}
