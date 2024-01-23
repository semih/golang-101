package main

import "fmt"

func main() {

	/*var a int
	// Create a pointer with type of int
	var pointer *int
	a = 10

	// Assign the address value(&) of a to pointer
	pointer = &a

	fmt.Println(a)       // -> 10
	fmt.Println(pointer) // -> 0x14000112018

	// Get the value from the address with *
	fmt.Println(*pointer)

	// Update the value of the pointer
	*pointer = 20
	fmt.Println(a)
	*/

	/*
		var x = 10
		var y int
		var p *int

		y = x
		p = &x

		*p = 20
		fmt.Println(x, y) // 20 10
	*/

	/*var x int = 10

	add12(x)
	fmt.Println(x)

	add12pointer(&x)
	fmt.Println(x)
	*/

	var numbers = []int{1, 2, 3}
	changeValue(numbers) // -> [1000 2 3]
	fmt.Println(numbers)
}

func changeValue(numbers []int) {
	numbers[0] = 1000
}

func add12(x int) { // pass by value
	x = x + 12
}

func add12pointer(p *int) { // pass by reference
	*p = *p + 12
}
