package main

import "fmt"

type myStruct struct {
	FirstName string
}

// function receiver
func (m *myStruct) returnMyStruct() string {
	return m.FirstName
}

func main() {
	user1 := myStruct{
		FirstName: "John",
	}

	user2 := myStruct{
		FirstName: "Mary",
	}

	fmt.Println(user1.returnMyStruct())
	fmt.Println(user2.returnMyStruct())
}
