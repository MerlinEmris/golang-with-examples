package main

import "fmt"

type EngiPer struct {
	Name string
}

func (e *EngiPer) UpdateName(name string) {
	e.Name = name
}

func doStuff(e *EngiPer) {
	name := "Gurban"
	defer e.UpdateName(name)
	name = "Maksat"
}

func main() {
	elli := &EngiPer{
		Name: "Meret",
	}
	doStuff(elli)
	fmt.Printf("%+v\n", elli)
}
