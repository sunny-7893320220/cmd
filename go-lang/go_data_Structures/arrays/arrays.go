/*

The array is the collection similar type of data


for the explanation of the array we can use the below example

I want to stroage the 5 varibale. By creating the 5 different variables it will take the 5 memory locations.

Simply, I can create the array of size 5 and store the 5 varibale.

--> The array has the fix size of storage.

synatx of array

var array_variable = [size]datatype{elements of array}


NOTE:
1. The array has the fix size of storage.
2. While printing the array, it is printing in this way []
3. we can create the array without fix size.

syntax of unfix size array

var array_variable = [...]datatype{elements of array}
*/

package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr[2])
	fmt.Println(arr[4:7])

	arr2 := [2][2]int{{1, 2}, {3, 4}}

	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Println(arr2[i][j])
		}
	}

	fmt.Println(arr2[1][1])
}
