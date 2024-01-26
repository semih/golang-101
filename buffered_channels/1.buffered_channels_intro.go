package main

import (
	"fmt"
)

func main() {
	// https://go.dev/src/runtime/chan.go
	// you have to set size to create a buffered channel
	myChannel := make(chan int, 2)

	//myChannel <- 1
	//myChannel <- 2
	//myChannel <- 3 // Blocking - channel is full so, it gives deadlock error

	/*data := <-myChannel // it gives an error since the channel is empty
	fmt.Println(data)*/

	myChannel <- 1
	myChannel <- 2
	fmt.Println(<-myChannel)
	myChannel <- 3
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)

}
