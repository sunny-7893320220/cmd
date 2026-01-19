/*

We can also loop through the slice using the for loop.


*/

package main

import "fmt"

func main() {
	var sclie = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(sclie); i++ {
		fmt.Println(sclie[i])
	}
}
