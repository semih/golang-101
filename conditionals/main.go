package main

import "fmt"

func main() {

	/*var age = 17
	if age >= 18 {
		fmt.Println("You can vote!")
	} else {
		fmt.Println("You can not vote!")
	}*/

	a := 10
	b := 100
	c := 30

	if a >= b && a >= c {
		fmt.Println("a is the greatest one")
	} else if b >= a && b >= c {
		fmt.Println("b is the greatest one")
	} else {
		fmt.Println("c is the greatest one")
	}
}
