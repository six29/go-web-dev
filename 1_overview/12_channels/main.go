package main

import (
	"log"

	"github.com/six29/go-web-dev/helpers"
)

var numPool = 10

func CalVal(intChan chan int) {
	randomNum := helpers.RandomNumber(numPool)
	intChan <- randomNum
}

func main() {
	intChannel := make(chan int)
	defer close(intChannel)

	go CalVal(intChannel)

	num := <-intChannel

	log.Println("Value of intChannel", num)
}
