package main

import (
	"fmt"
)

func main() {
	//f1()
	/*func() {
		fmt.Println("f1")
	}()

	total := func(x int, y int) int {
		return x + y
	}(3, 5)
	fmt.Println(total)
	*/

	add := func(x int, y int) int {
		return x + y
	}

	//fmt.Println(reflect.TypeOf(add))
	//fmt.Println(add(3, 5))

	multiply := func(x int, y int) int {
		return x * y
	}

	a, b := calculator(3, 5, add, multiply)
	fmt.Println(a, b)
}

func f1() {
	fmt.Println("f1")
}

func calculator(x int, y int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
