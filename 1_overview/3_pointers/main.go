package main

import "fmt"

func main() {
	var myString string
	myString = "Green"
	fmt.Println(myString)
	changeUsingPointer(&myString)
	fmt.Println(myString)
}

func changeUsingPointer(s *string) {
	fmt.Println(s) // returns the memory address of myString
	newValue := "Red"
	*s = newValue // put newValue to the memory address
}
