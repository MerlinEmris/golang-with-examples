package main

import (
	"fmt"
	"time"
)

func SayHello(name string) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Hello %s\n", name)
}

func main() {
	go SayHello("Merlin")
	fmt.Println("It is printing first")
	time.Sleep(3 * time.Second)
}
