package main

import (
	"errors"
	"fmt"
)

func main() {
	mgs := "Hello"
	fmt.Println(mgs)

	error := errors.New("Something went wrong")
	if mgs != "hello_world" {
		fmt.Println(error)
	}

}
