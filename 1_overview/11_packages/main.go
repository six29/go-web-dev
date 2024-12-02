package main

import (
	"fmt"

	"github.com/six29/go-web-dev/helpers"
)

func main() {
	var myType helpers.MyType
	myType.TypeName = "Tyson"
	myType.TypeNumber = 50
	fmt.Println(myType.TypeName)
	fmt.Println(myType.TypeNumber)
}
