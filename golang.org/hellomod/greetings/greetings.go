package greetings

import (
	"fmt"
	"math/rand"
	"time"
	"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randVal(), name)
	return message,nil
}

//func Init() {
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randVal() string {
	rs := []string {
		"Hi %v, welcome!",
		"Great to see you, %v",
		"Yo %v!",
	}
	return rs[rand.Intn(len(rs))]

}
