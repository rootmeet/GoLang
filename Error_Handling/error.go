package main

import (
	"errors"
	"fmt"
)

func main() {
	message := "Programming"

	myError := errors.New("Wrong Message")

	if message == "Programming" {
		fmt.Println(message)
	} else {
		fmt.Println(myError)
	}
}
