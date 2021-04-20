package main

import "log"

func main() {
	var whatToSay string

	// _ means the second param, but we do not need it
	whatToSay, _ = saySomething("Hi")

	log.Println(whatToSay)

	changeUsingPointer(&whatToSay)

	log.Println("after func call what to say is set to", whatToSay)
}

func saySomething(say string) (string, string) {
	return say, "world"
}

func changeUsingPointer(say *string) {
	newValue := "Red"
	*say = newValue
}
