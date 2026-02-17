package main

import "fmt"

func main() {
	defer fmt.Println("One")
	fmt.Println("Two")
	fmt.Println("Three")
	defer fmt.Println("Four")
}
