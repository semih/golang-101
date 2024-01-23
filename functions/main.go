package main

import "fmt"

func main() {
	/*total, difference, multiply := calculation(10, 20)
	fmt.Printf("Total %d, Difference %d, Multiply %d", total, difference, multiply)*/

	/*var numbers = []int{10, 20, 2, 3, 5}
	fmt.Println(sum(numbers))*/

	fmt.Println(sum(10, 20, 2))
}

/*func add(a int, b int) {
	fmt.Println(a + b)
}*/

/*func add(a int, b int) int {
	return a + b
}*/

/*func calculation(a int, b int) (int, int, int) {
	return a + b, a - b, a * b
}*/

/*func sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}*/

/*func sum(a int, b int, c int) int {
	return a + b + c
}*/

func sum(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
