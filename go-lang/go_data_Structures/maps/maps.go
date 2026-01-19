/*

In go, the maps store the elements in the key/value pairs.


Syntax

mapName := map[keyType]valueType{
	key1: value1,
	key2: value2,
	...
}

*/

package main

import "fmt"

func main() {
	var map1 = map[string]int{}
	// fmt.Println(len(map1))
	map1 = map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(map1["one"])
	fmt.Println(map1["two"])
}
