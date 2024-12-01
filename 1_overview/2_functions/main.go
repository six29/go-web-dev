package main

import (
	"fmt"
	"log"
)

func main() {
	var whatToSay string
	var saySomethingElse string
	whatToSay = returnSomething("Hello")
	log.Println(whatToSay)
	saySomethingElse = "Hi, there!"
	log.Println(saySomethingElse)

	fmt.Println(returnSomething("How are you?"))

	fmt.Println(returnTwoValues("You are my only ", 1))

	var s string
	var n int
	s, n = returnTwoValues("You are my only ", 1)
	fmt.Println(s, n)

	var x string
	x, _ = returnTwoValues("You are my only ", 1)
	log.Println(x)
}

func returnSomething(s string) string {
	return s
}

func returnTwoValues(s string, n int) (string, int) {
	return s, n
}
