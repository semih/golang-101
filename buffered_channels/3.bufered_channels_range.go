package main

import (
	"fmt"
	"time"
)

func main() {
	bufferedChannel := make(chan int, 4)

	go func() {
		for i := 1; i <= 10; i++ {
			bufferedChannel <- i
			fmt.Println("Send data: ", i)
		}
		close(bufferedChannel)
	}()

	time.Sleep(3 * time.Second)
	for data := range bufferedChannel {
		fmt.Println("Received data: ", data)
		time.Sleep(1 * time.Second)
	}

}
