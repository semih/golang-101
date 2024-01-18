package main

import (
	"fmt"
)

func main() {
	/*
		var productName string
		var quantity int
		var discount float32
		var isInStock bool

		productName = "golang"
		quantity = 10
		discount = 0.37
		isInStock = false

		fmt.Println(productName, reflect.TypeOf(productName))
		fmt.Println(quantity, reflect.TypeOf(quantity))
		fmt.Println(discount, reflect.TypeOf(discount))
		fmt.Println(isInStock, reflect.TypeOf(isInStock))
	*/

	/*var productName string = "golang"
	fmt.Println(productName, reflect.TypeOf(productName))*/

	/*
		productName := "golang"
		fmt.Println(productName, reflect.TypeOf(productName))

		quantity := 10
		fmt.Println(quantity, reflect.TypeOf(quantity))

		discount := 0.2
		fmt.Println(discount, reflect.TypeOf(discount))*/

	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "golang"
	quantity = 10
	discount = 0.37
	isInStock = true

	/*fmt.Println("Product name: ", productName, "Quantity: ", quantity, "Discount: ", discount, "IsInStock: ", isInStock)*/
	/*fmt.Printf("Product name: %s, Quantity: %d, Discount: %f, IsInStock: %t\n", productName, quantity, discount, isInStock)*/
	fmt.Printf("Product name: %v, Quantity: %v, Discount: %v, IsInStock: %v\n", productName, quantity, discount, isInStock)
}
