package main

import "fmt"

func main() {
	var a int 
	var b int
	fmt.Println("Enter a number")
	fmt.Scan(&a)
	fmt.Println("Enter a number")
	fmt.Scan(&b)
	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("b is greater than a")
	}
}