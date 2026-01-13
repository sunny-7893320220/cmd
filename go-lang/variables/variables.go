// the variables are like the containers that is used to store the data.

/*

   they are 3 types to determine it.

   var a int = 10


   'a' is the name of the variable which is an integer

   var a = 10

   a := 10  ==> this is called the "shorthand notation"


   const l = 30 ==> if we use this const we can not change this and we can't use the := sambol for the const

*/

package main

import (
	"fmt"
)

func main() {
	var a int = 20

	fmt.Println(a)

	var b = 30

	fmt.Println(b)

	c := 40

	fmt.Println("Initial number value", c)

	c = 50

	fmt.Println("Updated number value", c)

	var name, age = "saikrishna", 22

	fmt.Println(name, age)

	name, age = "sunny", 32

	fmt.Println("changed name and age is", name, age)

	const l int = 30

	fmt.Println("initial value of const to test", l)

	var mgs string

	mgs = "welcome to go language"

	fmt.Println(mgs)

	message := "welcome to go language for secoud time"

	fmt.Println("this is the message i want to display", message)

	company := "deepta"

	var position = "developer"

	fmt.Println("my company is", company, "and my position is", position)

	fmt.Println("company is %s and the position is %s", company, position)

	var d string
	var e int

	fmt.Print("Enter you name: ")
	fmt.Scan(&d)

	fmt.Print("Enter you age: ")
	fmt.Scan(&e)

	fmt.Println("your name is", d)
	fmt.Println("your age is", e)

}
