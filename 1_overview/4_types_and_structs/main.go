package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Paul",
		LastName:  "Figures",
	}
	fmt.Println(user.FirstName, user.LastName, user.BirthDate)
}
