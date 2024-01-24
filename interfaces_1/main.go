package main

import "fmt"

type Shippable interface {
	ShippingCost() int
}

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

type Flower struct {
	desi int
}

func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}

func (electronic *Electronic) ShippingCost() int {
	return 10 + electronic.desi*3
}

func main() {
	/*book1 := &Book{desi: 10}
	book2 := &Book{desi: 20}
	fmt.Println(book1.ShippingCost())
	fmt.Println(book2.ShippingCost())*/

	/*electronic1 := &Electronic{desi: 20}
	fmt.Println(electronic1.ShippingCost())*/

	/*books := []Book{{desi: 10}, {desi: 20}}
	fmt.Println(calculatetotalShippingCostOfBooks(books))*/

	/*electronics := []Electronic{{desi: 10}, {desi: 20}}
	totalCost := calculatetotalShippingCostOfElectronics(electronics)
	fmt.Printf("Electronic shipping cost: %d\n", totalCost)*/

	/*var product Shippable
	product = &Book{desi: 10}
	fmt.Println(product.ShippingCost())

	product = &Electronic{desi: 10}
	fmt.Println(product.ShippingCost())*/

	var products []Shippable = []Shippable{&Book{desi: 10}, &Book{desi: 20}, &Electronic{desi: 20}}
	fmt.Println(calculatetotalShippingCost(products))
}

/*
func calculatetotalShippingCostOfBooks(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.ShippingCost()
	}
	return total
}

func calculatetotalShippingCostOfElectronics(electronics []Electronic) int {
	total := 0
	for _, electronic := range electronics {
		total += electronic.ShippingCost()
	}
	return total
}
*/

func calculatetotalShippingCost(products []Shippable) int {
	total := 0
	for _, product := range products {
		total += product.ShippingCost()
	}
	return total
}
