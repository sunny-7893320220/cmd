/*

A string is a sequence of characters

Functions	Descriptions
Compare()	compares two strings
Contains()	checks if a substring is present inside a string
Replaces()	replaces a substring with another substring
ToLower()	converts a string to lowercase
ToUpper()	converts a string to uppercase
Split()	    splits a string into multiple substrings

*/

package main

import "fmt"

func main() {
	var str string = `Hello`
	fmt.Println(str)
	fmt.Println(len(str))
	fmt.Println(string(str[0]))
	fmt.Println(string(str[1]))
	fmt.Println(string(str[2]))
	fmt.Println(string(str[3]))
	fmt.Println(string(str[4]))
}
