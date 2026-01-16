/*
In go, the switch statement is used to perform different actions based on different conditions.

In Go, the switch statement allows us to execute one code block among many alternatives.

refer to the photo in the same directory to understand the switch statement

*/

package main

import "fmt"

func main() {
	var num int
	fmt.Println("what is the week number")
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid week number")
	}
}