package main

import "fmt"

func main() {
	// Go routines communicate with each other through channels.
	/*
		data := go f1() // -> syntax error
		fmt.Println(data)
	*/

	// if you don't set a size, then it will be created an unbuffered channel.
	myChannel := make(chan string)
	done := make(chan bool)
	go func() {
		myChannel <- "Hello World"
	}()
	// Blocking - if there is no message in channel, the program will be blocked.
	// isOpen: channel is open
	/*
		message, isOpen := <-myChannel
		fmt.Println(message, isOpen)
	*/

	anotherChannel := make(chan string)
	go func() {
		message := <-anotherChannel
		fmt.Println(message)
		done <- true
	}()

	go func() {
		anotherChannel <- "Hello World"
	}()

	<-done
	fmt.Println("End of the main")
}

/*func f1() int {
	return 1
}*/
