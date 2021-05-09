package main

import (
	"fmt"
	"log"
	"tuts/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	if (err!= nil) {
		log.Fatal(err)
	}
	message,err = greetings.Hello("Ruby H")
	fmt.Println(message)
}
