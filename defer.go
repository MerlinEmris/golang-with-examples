package main

import (
	"errors"
	"fmt"
)

func SomeFunc() (err error) {
	defer func() {
		err = errors.New("some error")
	}()
	return nil
}

func returnValue() (x int) {
	defer func() {
		x = 10
	}()
	x = 5
	return
}

func main() {

	if err := SomeFunc(); err != nil {
		fmt.Println(err.Error())
	}

	defer func() {
		fmt.Println("defer in main")
	}()
	defer func() {
		fmt.Println("this one runs first")
	}()
	fmt.Println("some logic")
	fmt.Println(returnValue())
}
