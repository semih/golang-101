package main

import "fmt"

func main() {
	//var numbers = []int{1, 2, 3, 4}

	/*for index:=0; index<len(numbers); index++ {
		fmt.Println(numbers[index])
	}*/

	/*for index, value := range numbers {
		fmt.Println(index, value)
	}*/

	/*for _, value := range numbers {
		fmt.Println(value)
	}*/

	/*var language = "Golang"
	for _, character := range language {
		fmt.Println(character)
	}

	for _, character := range language {
		fmt.Printf("Character %c\n", character)
	}*/

	var animals = map[string]int{
		"Cat":  1,
		"Dog":  2,
		"Bird": 3,
	}

	for key, value := range animals {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

}
