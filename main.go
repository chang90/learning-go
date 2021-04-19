package main

import "log"

func main() {
	var whatToSay string

	// _ means the second param, but we do not need it
	whatToSay, _ = saySomething("Hi")

	log.Println(whatToSay)
}

func saySomething(say string) (string, string) {
	return say, "world"
}
