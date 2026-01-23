/*

we use the funcions to divide the code into smaller chunks to make our code looks clean and easy to read

syntax

func greet(){
  // code
}

func function_name(Parameter-list)(Return_type){
    // function body.....
}

& --> memory address of value

* --> address of value

*/

package main

import "fmt"

// func add(num1, num2 *int) int {
// 	fmt.Printf("num1{%p} num2{%p} address from add function:\n", num1, num2)
// 	return *num1 + *num2
// }

// func main() {
// 	num1 := 1
// 	num2 := 2
// 	fmt.Printf("num1{%p} num2{%p} address from main function:\n", &num1, &num2)
// 	result := add(&num1, &num2)
// 	fmt.Println(result)
// }

func add(num1, num2 *int) (int, int) {
	sum := *num1 + *num2
	difference := *num1 - *num2
	return sum, difference
}

func main() {
	num1 := 1
	num2 := 2
	sum1, difference1 := add(&num1, &num2)
	fmt.Println(&num1)
	fmt.Println(&num2)
	fmt.Println(sum1, difference1)
}
