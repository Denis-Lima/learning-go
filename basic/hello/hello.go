package main

import (
	"fmt"
	"hello/basic/greetings"
)

func main() {
	message := greetings.Hello("Denis")
	fmt.Println(message)
}