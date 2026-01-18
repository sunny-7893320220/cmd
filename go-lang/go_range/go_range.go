/*

in the go language the range keyword is used to iterate over a collection of values.

like : "array", "slice", "map", "string"


mostly we use the range keyword with the for loop.
*/

package main

import (
	"fmt"
)

func main() {

	/*

		This is the example of the range keyword with the array.

	*/

	// var arr = [5]int{1, 2, 3, 4, 5}

	// for i, v := range arr {
	// 	fmt.Println("numbers", i, "=", v)
	// }

	/*

		This is the example of the range keyword with the string.

	*/
	// string1 := "golang"

	// for i, s := range string1 {
	// 	fmt.Println(i, string(s))
	// }

	// string := "Golang"
	// fmt.Println("Index: Character")

	// // i access index of each character
	// // item access each character
	// for i, item := range string {
	// 	fmt.Printf("%d= %c \n", i, item)
	// }

	/*

		This is the example of the range keyword with the maps.

	*/

	map1 := map[string]float32{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	for k := range map1 {
		fmt.Println(k)
	}

	for i := 1; i <= 10; i++ {

		for j := 10; j <= 20; j++ {
			if i == j {
				fmt.Println(i)
				break
			}
		}

	}

}
