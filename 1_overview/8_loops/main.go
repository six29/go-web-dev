package main

import "log"

func main() {
	mySlice := []string{"Cat", "Mouse", "Dog", "Banana"}

	for key, val := range mySlice {
		log.Println(key, val)
	}

	myMap := make(map[string]string)
	myMap["Pet"] = "Dragon"
	myMap["Job"] = "Knight"
	myMap["Skill"] = "Eating"
	myMap["NetWorth"] = "Ten billion coins"

	for key, val := range myMap {
		log.Println(key, val)
	}
}
