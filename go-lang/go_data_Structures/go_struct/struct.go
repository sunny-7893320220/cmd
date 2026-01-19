package main

import "fmt"

/*
Go struct is a collection of fields.

A struct is used to store the variable of different data types.

==> struct is a blueprint. To use a struct, we need to create an instance of it.

# SYNTAX

	type structName struct {
		fieldName1 fieldType1
	}
*/

// Person is the "Blueprint" (Type Definition)
type Person struct {
	Name string
	Age  int
}

type Rectangle struct {
	Length int
	Width  int
}

func main() {
	// Creating an "Instance" of Person
	// 'p1' is a concrete object created from the Person blueprint

	// This is the one method of creating an instance of a struct

	// var p0 Person
	// p0.Name = "John"
	// p0.Age = 30
	// fmt.Println(p0)

	// var p1 Person
	// p1.Name = "John"
	// p1.Age = 30
	// fmt.Println(p1)

	// var p2 Person
	// p2.Name = "John"
	// p2.Age = 30
	// fmt.Println(p2)

	// This is the second method of creating an instance of a struct

	po := Person{Name: "John", Age: 30}
	fmt.Println(po)

	p1 := Person{Name: "John", Age: 31}
	fmt.Println(p1)

	p2 := Person{Name: "John", Age: 32}
	fmt.Println(p2)

	r1 := Rectangle{Length: 10, Width: 20}
	fmt.Println(r1)
	fmt.Println(r1.Length)
	fmt.Println(r1.Width)
	fmt.Println(r1.Length * r1.Width)

	r2 := Rectangle{Length: 11, Width: 21}
	fmt.Println(r2)
	fmt.Println(r2.Length)
	fmt.Println(r2.Width)
	fmt.Println(r2.Length * r2.Width)

	r3 := Rectangle{Length: 12, Width: 22}
	fmt.Println(r3)
	fmt.Println(r3.Length)
	fmt.Println(r3.Width)
	fmt.Println(r3.Length * r3.Width)
}
