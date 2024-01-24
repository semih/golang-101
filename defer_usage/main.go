package main

import "fmt"

func main() {
	/*defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")

	fmt.Println("Hello World")*/

	var condition = true
	defer cleanup()

	if condition {
		panic("An error occurred!")
	}
}

func cleanup() {
	fmt.Println("Cleanup worked...")
}

// the meaning of Defer is to postpone.
// deferred function calls are pushed onto a STACK.
// when a function returns, its deferred calls are executed in last-in-first-out order

// it should be used like a finally statements.
// close the file, close the channel
