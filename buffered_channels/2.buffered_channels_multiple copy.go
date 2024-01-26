package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)
	go func() {
		data1 := <-messages
		fmt.Println("First: ", data1)
		data2 := <-messages
		fmt.Println("Second: ", data2)
	}()

	messages <- "Hello"
	messages <- "World"
	messages <- "!"
	time.Sleep(1 * time.Second)
	fmt.Println("End of the main")
}
