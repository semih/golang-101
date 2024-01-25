package main

import "fmt"

var x = 10

func main() {

	var condition = true
	if condition {
		// block scope
		var x = 10
		fmt.Println(x)
	}

	// fmt.Println(x) -> compiler gives error (block scope)

	test()
	// fmt.Println(x, y) -> compiler gives error (function scope)

	fmt.Println(x) // -> doesn't give error (global scope)
}

func test() {
	var x = 30 // function scope
	fmt.Println(x)
}
