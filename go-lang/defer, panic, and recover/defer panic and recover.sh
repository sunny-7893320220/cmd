/* In Go, we use defer, panic and recover statements to handle errors.

DEFER

we use the defer statement to delay the execution of a function that migth cause an error.

PANIC

The panic statement terminates the program immediately.

RECOVER

The recover is use to recover the mgs during the panic


*/


we use the defer statement to prevent the exection of a function until all the other functions are executed.

package main
import "fmt"

func main() {

  // defer the execution of Println() function
  defer fmt.Println("Three")

  fmt.Println("One")
  fmt.Println("Two")

}


Golang panic

We use the panic statement to immediately end of exection of a program.If our program reaches a point where it cannot be recovered due to some major errors, it's best to use panic.


package main
import "fmt"

func main() {

  fmt.Println("Help! Something bad is happening.")
  panic ("Ending the program")
  fmt.Println("Waiting to execute")

}