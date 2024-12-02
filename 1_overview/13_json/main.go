package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDragon bool   `json:"has_dragon"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Spider",
		"last_name": "Man",
		"hair_color": "red",
		"has_dragon": true
	},
	{
		"first_name": "Xy",
		"last_name": "Clops",
		"hair_color": "blue",
		"has_dragon": false
	}
]`

	// Reading json to a struct
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmashalling json", err)
	}

	log.Printf("Unmarshalled %v", unmarshalled)

	// writing struct to a json
	var mySlice []Person

	var m1 Person

	m1.FirstName = "Jean"
	m1.LastName = "Gray"
	m1.HairColor = "Red"
	m1.HasDragon = false

	mySlice = append(mySlice, m1)

	var m2 Person

	m2.FirstName = "Logan"
	m2.LastName = "Wolf"
	m2.HairColor = "Yellow"
	m2.HasDragon = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(string(newJson))
}
