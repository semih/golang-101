package main

import (
	"context"
	"fmt"
	"time"
)

// to manage the lifecycle and control the cancellation of operations within a concurrent or asynchronous environment.
type Product struct {
	Id   int64
	Name string
}

var productChannel = make(chan Product)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go getProductDetailsFromExternalService(114231)
	select {
	case productDetail := <-productChannel:
		fmt.Println("Received product details: ", productDetail)
	case <-ctx.Done():
		fmt.Println("Timeout occurred, context canceled.")
	}
	fmt.Println("End of the main")
}

func getProductDetailsFromExternalService(productId int64) {
	time.Sleep(10 * time.Second) // Timeout occurred, context canceled.
	//time.Sleep(2 * time.Second) // Received product details:  {114231 Shoes}

	productChannel <- Product{
		Id:   productId,
		Name: "Shoes",
	}
}
