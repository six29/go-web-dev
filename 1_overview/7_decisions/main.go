package main

import "fmt"

func main() {
	var isTrue bool
	isTrue = true
	myNum := 200

	// if
	if isTrue && myNum < 99 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	// switch
	// will automatically break when a match is found
	switch myNum {
	case 80:
		fmt.Println("Eighty")
	case 90:
		fmt.Println("Ninety")
	default:
		fmt.Println("Whatever")
	}
}
