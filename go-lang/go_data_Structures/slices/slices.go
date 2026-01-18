/*
The slice is the dynamic size of storage.

It does not have the fix size of storage.

The slice is a collection of similar type of data, just like array.

However, unlike arrays, slice doesn't have a fixed size. We can add or remove elements from the array.

syntax

sliceName := []type{value1, value2, ...}

In the sclie we have to 4 in-build functions.

1.append() --> adds element to a slice
2.copy() --> copy elements of one slice to another
3.Equal() --> compares two slices
4.len() --> find the length of a slice
*/
package main

import (
	"fmt"
)

func slice_main() {

	var sclie = []int{1, 2, 3, 4, 5}
	fmt.Println(sclie)
	sclie = append(sclie, 6, 7, 8, 9, 10)
	fmt.Println(sclie)
	// and we can add the one slice to another slice
	sclie2 := []int{11, 12, 13, 14, 15}
	sclie = append(sclie, sclie2...)
	fmt.Println(sclie)
	// and we can copy the one slice to another slice
	sclie3 := make([]int, len(sclie))
	copy(sclie3, sclie)
	fmt.Println(sclie3)

}
