package main

import "fmt"

func main() {
	//Fixed Arrays
	/*var letters [3]string
	letters[0] = "abc"
	letters[1] = "def"
	letters[2] = "ghi"
	fmt.Println(letters)*/

	/*var letters = [3]string{"abc", "def", "ghi"}
	fmt.Println(letters)
	fmt.Println(letters[0:2])*/

	//Slices
	var letters = []string{"abc", "def", "ghi", "jkl"}
	fmt.Println(letters)

	letters = append(letters, "mno")
	fmt.Println(letters)
}
