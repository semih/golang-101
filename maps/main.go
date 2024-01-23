package main

import "fmt"

func main() {
	/*var animals map[string]int
	animals = make(map[string]int, 0)

	animals["Cat"] = 1
	animals["Dog"] = 2
	animals["Bird"] = 3

	fmt.Println(animals)
	fmt.Println(animals["Cat"])  // -> 1
	fmt.Println(animals["Bear"]) // -> 0
	*/

	animals := map[string]int{
		"Cat":  1,
		"Dog":  2,
		"Bird": 3,
	}
	fmt.Println(animals)
	delete(animals, "Cat")
	fmt.Println(animals)
}
