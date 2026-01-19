/*

The string we can compare two string .

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "Hello"
	var str2 string = "Hello"
	var str3 string = "go"
	fmt.Println(strings.Compare(str1, str2))
	result := strings.Contains(str1, str3)
	fmt.Println(result)
	result = strings.Contains(str1, str2)
	fmt.Println(result)

	result1 := strings.Replace(str1, "Hello", "World", 1)
	fmt.Println(result1)

	result2 := strings.ToLower(str1)
	fmt.Println(result2)

	result3 := strings.ToUpper(str1)
	fmt.Println(result3)

}
