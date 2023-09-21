package main

import "fmt"

// Engineer - stores the name and age of an engineer
type Engi struct {
	Name string
	Age  int
}

func (e *Engi) UpdateAge() {
	e.Age += 1
}

func (e *Engi) UpdateName() {
	e.Name = "Marry"
}
func UpdateAge(e *Engi) {
	e.Age += 1
}

func main() {
	var name = "mary"
	fmt.Println(name)

	ptr := &name
	fmt.Println(ptr)

	*ptr = "tkm"
	fmt.Println(name)

	mary := &Engi{
		Name: "Mary",
		Age:  33,
	}
	mary.UpdateAge()
	mary.UpdateName()
	fmt.Println(mary)

	UpdateAge(mary)
	fmt.Println(mary)
}
