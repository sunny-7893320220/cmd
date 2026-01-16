/* This is the for loop syntax in the go language 

for initialization; condition; update {
  statement(s)
}

*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Capture start time
	start := time.Now()

	// Print initial memory usage
	printMemUsage()

	// for i := 0; i < 100; i++ {
	// 	// Note: Println does not support %d formatting, use Printf for that.
	// 	// However, leaving as is to match your style.
	// 	fmt.Println("This is the loop of %d", i)
	// }

	// Print final memory usage
	// printMemUsage()

	var n, sum = 10000000, 0

	for i := 1; i <= n ; i++ {
		fmt.Println(i , sum)
		sum += i
	}
	fmt.Println("The sum of first ", n, " natural numbers is ", sum)
	
	// Print execution time
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)

	printMemUsage()
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("\nAlloc = %v KiB", m.Alloc/1024)
	fmt.Printf("\tTotalAlloc = %v KiB", m.TotalAlloc/1024)
	fmt.Printf("\tSys = %v KiB", m.Sys/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}