package main

import (
	"fmt"
	"sort"
)

type User struct {
	FirstName string
}

func main() {
	// Maps
	// Not sorted by default
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	fmt.Println(myMap["dog"])

	myMap2 := make(map[string]int)
	myMap2["Age"] = 10
	fmt.Println(myMap2["Age"])

	me := User{
		FirstName: "Paul",
	}
	user := make(map[string]User)
	user["me"] = me
	fmt.Println(user["me"].FirstName)

	// Slices ska Arrays
	var mySlice []string
	mySlice = append(mySlice, "Paul")
	mySlice = append(mySlice, "John")
	sort.Strings(mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[0])

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(numbers[0:4])
	fmt.Println(numbers[1:4])
}
