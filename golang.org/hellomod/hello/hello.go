package main

import (
	"fmt"
	"tuts/greetings"
)

func main() {
	message := greetings.Hello("Ruby H")
	fmt.Println(message)
}
