package main

import (
	"errors"
	"fmt"
)

func main() {
	message := "Hello Error"

	myError := errors.New("Wrong Message")

	if message != "Programming" {
		fmt.Println(myError)
	}
}
