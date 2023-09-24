package main

import (
	"errors"
	"fmt"
)

func letspunic() {
	defer func() {
		fmt.Println("I am running in the end")
	}()
	panic("panic <><><>")
}

func someFunction() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic")
			err = errors.New("error:=>some panic happened")
		}
	}()
	letspunic()
	return nil
}

func main() {
	defer func() {
		fmt.Println("end of main func")
	}()
	if err := someFunction(); err != nil {
		fmt.Println("error happened")
		fmt.Println(err.Error())
	}
}
