package main

import (
	"fmt"
)

func main() {
	/*var x int
	var y float32
	var pointer1 *int

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(pointer1) // nil

	if pointer1 == nil {
		fmt.Println("Pointer is nil")
	}
	*/

	/*fileData, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileData)
	}
	*/

	/*result, err := divide(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}*/

	productService := ProductService{}

	err := productService.Add(Product{id: 1, name: "Shoes", price: 3000})
	if err != nil {
		fmt.Println(err)
	}
}

/*func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Division by zero error!")
	}
	return x / y, nil
}*/

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct {
}

func (productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		//return errors.New("Product name cannot be null!")
		return &ValidationError{text: "Product name should not be null!", code: "1001"}
	}
	if product.price <= 0 {
		//return errors.New("Product price should be greater than zero!")
		return &ValidationError{text: "Product price must be greater than zero!", code: "1002"}
	}
	fmt.Println("Product has been added.")
	return nil
}

type ValidationError struct {
	code string
	text string
}

func (validationError *ValidationError) Error() string {
	return fmt.Sprintf("Error Message: %s , Error Code: %s", validationError.text, validationError.code)
}
