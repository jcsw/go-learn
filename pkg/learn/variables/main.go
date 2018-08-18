package main

import "fmt"

var (
	firstName, lastName string
	age                 int
	foreign             bool
	weinght             float64
	address             map[string]string
)

func main() {
	fmt.Printf("firstName: %s\r\n", firstName)
	fmt.Printf("lastName: %s\r\n", lastName)
	fmt.Printf("age: %d\r\n", age)
	fmt.Printf("foreign: %v\r\n", foreign)
	fmt.Printf("weinght: %f\r\n", weinght)
	fmt.Printf("address: %v\r\n", address)

	fmt.Printf("\n")

	firstName = "Robert"
	lastName = "Martin"
	age = 57
	foreign = true
	weinght = 64.35
	address = make(map[string]string)
	address["home"] = "Street 54, New York, number 1255"
	address["job"] = "Street 85, New York, number 19"

	fmt.Printf("firstName: %s\r\n", firstName)
	fmt.Printf("lastName: %s\r\n", lastName)
	fmt.Printf("age: %d\r\n", age)
	fmt.Printf("foreign: %v\r\n", foreign)
	fmt.Printf("weinght: %f\r\n", weinght)

	fmt.Printf("address size: %d\r\n", len(address))
	for key, value := range address {
		fmt.Printf("address: %s > %s\r\n", key, value)
	}

	delete(address, "job")
	fmt.Printf("address size: %d\r\n", len(address))
	for key, value := range address {
		fmt.Printf("address: %s > %s\r\n", key, value)
	}
}
