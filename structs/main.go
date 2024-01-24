package main

import "fmt"

func main() {
	var customer1 = Customer{id: 1, name: "John", age: 30, address: Address{code: 867, city: "New York"}}
	/*var customer2 = Customer{id: 2, name: "Jack", age: 35, address: Address{code: 231, city: "Houston"}}*/

	/*customer1.age = 38
	fmt.Println(customer1)
	fmt.Println(customer2)*/
	// changeName(customer1) -> it updates the value, not reference.
	customer1.ToString()

	changeNameWithReference(&customer1)
	customer1.ToString()

	customer1.changeName("James") // -> it updates the value, not reference.
	customer1.ToString()

	customer1.changeNameWithPointer("James") // -> it updates the reference.
	customer1.ToString()
}

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}
type Address struct {
	code int
	city string
}

func (customer Customer) ToString() {
	fmt.Printf("Name: %s , Age: %d\n", customer.name, customer.age)
}

/*func changeName(customer Customer) {
	customer.name = "Sawyer"
}*/

func changeNameWithReference(customer *Customer) {
	customer.name = "Sawyer"
}

func (customer Customer) changeName(newName string) {
	customer.name = newName
}

func (customer *Customer) changeNameWithPointer(newName string) {
	customer.name = newName
}
