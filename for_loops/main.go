package main

import "fmt"

func main() {

	/*index := 1
	for index <= 10 {
		fmt.Println(index)
		index++
	}*/

	/*for index := 1; index <= 10; index++ {
		fmt.Println(index)
	}*/

	//var letters = [3]string{"abc", "def", "ghi"}
	/*index := 0
	for index < len(letters) {
		fmt.Println(letters[index])
		index++
	}*/

	/*
		for index := 0; index < len(letters); index++ {
			fmt.Println(letters[index])
		}*/

	for index := 0; index <= 10; index++ {
		if index == 2 || index == 5 {
			continue
		}
		fmt.Println(index)
	}

}
