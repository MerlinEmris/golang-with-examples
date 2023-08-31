package main

import "fmt"

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

type Engineer struct {
	Name    string
	Age     int
	Project Project
}

func (e *Engineer) PrintData() {
	fmt.Println("Engineer Info:")
	fmt.Printf("%+v\n", e)
}

func (e *Engineer) SetAge(age int) {
	e.Age = age
	e.PrintData()
}

func main() {
	adam := Engineer{Name: "Jo", Age: 12, Project: Project{
		Name:         "Mary",
		Priority:     "High",
		Technologies: []string{"Vue", "GoLang"},
	}}
	fmt.Printf("%+v\n", adam)
	adam.SetAge(19)
}
