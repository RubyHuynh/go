package main

import (
	"fmt"
	"log"
	"tuts/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//greetings.Init()
	message, err := greetings.Hello("CatCat")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	pool := []string {"AA", "BB", "CC"}
	messages, err := greetings.Hellos(pool)
	fmt.Println(messages)
}
