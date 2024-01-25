package rest

import "fmt"

func Rest1() {
	rest2()
	fmt.Println("Rest1")
}

func rest2() { // -> It is not accessible from other packages
	fmt.Println("Rest2")
}
