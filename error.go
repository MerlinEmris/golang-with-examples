package main

import (
	"errors"
	"fmt"
)

var (
	ErrIsLessThanZero = errors.New("error:number less than zero")
	ErrIsNotEven      = errors.New("error:number is not even")
)

func isEven(num int) error {
	if num < 0 {
		return ErrIsLessThanZero
	} else if num%2 != 0 {
		return ErrIsNotEven
	}
	return nil
}

func main() {
	err := isEven(-23)
	if err != nil {
		if err == ErrIsNotEven {
			fmt.Println("Number is not even")
		} else if err == ErrIsLessThanZero {
			fmt.Println("Number is less than zero")
		}
		fmt.Println(err.Error())
	}
}
