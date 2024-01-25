package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//go f1() // -> goroutine might be defined go keyword
	//go f2()
	//fmt.Println("End of the main")
	//time.Sleep(1 * time.Second)

	/*
		// If we want all the go routines to be completed, we can use a wait group.
		wg := sync.WaitGroup{}
		wg.Add(2) // I will wait 2 goroutine to be completed

		go func() {
			defer wg.Done() // It should be done when function called.
			fmt.Println("f1")
		}()
		go func() {
			defer wg.Done()
			fmt.Println("f2")
		}()
		wg.Wait() // Blocked - if number of wg is not zero, will give fatal error: all goroutines are asleep - deadlock!
		fmt.Println("End of the main")
	*/

	startTime := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("f1")
		time.Sleep(3 * time.Second)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("f2")
		time.Sleep(3 * time.Second)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("f3")
		time.Sleep(3 * time.Second)
	}()
	wg.Wait()
	fmt.Println("Passed time:", time.Now().Sub(startTime))
}

/*func f1() {
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}*/
