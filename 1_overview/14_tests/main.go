package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(100.0, 0.0)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	fmt.Println("Result is", result)
}

func divide(x, y float32) (float32, error) {

	var result float32
	if y == 0 {
		return result, errors.New("Can not divide by 0")
	}
	result = x / y
	return result, nil

}
