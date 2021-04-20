package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	var whatToSay string

	// _ means the second param, but we do not need it
	whatToSay, _ = saySomething("Hi")

	log.Println(whatToSay)

	changeUsingPointer(&whatToSay)

	log.Println("after func call what to say is set to", whatToSay)

	log.Println(s)

	user := User{
		FirstName: "Emma",
		LastName:  "Sawyer",
	}

	log.Println(user.FirstName, user.LastName)

}

func saySomething(say string) (string, string) {
	return say, "world"
}

func changeUsingPointer(say *string) {
	newValue := "Red"
	*say = newValue
}
