/*
The converting of one data type to another data type is called type casting
we need to import the strconv package to convert int to string


--> we can convert int to float32 by using the float32(var)

--> In the same way we can convert float32 to int by using the int(var)

--> But we can't convert string to int directly we need to use the strconv package



*/


package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var a int = 10
	// var c float32 
	// fmt.Println("Enter a number")
	// fmt.Scan(&c)
	// var b float32 = float32(a) + (c)
	// fmt.Println(b)


	var intvar string = "My name is saikrishna"
	fmt.Println(intvar)

	var float32var int = 10
	fmt.Println(float32var)

	// Correct way to convert int to string
	var convertvar string = strconv.Itoa(float32var) + " " + intvar
	fmt.Println("The converted value is: ", convertvar)
}
