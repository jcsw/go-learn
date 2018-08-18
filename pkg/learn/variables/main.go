package main

import "fmt"

func main() {

	basics()

	maps()

	arrays()

	slices()
}

var (
	firstName, lastName string
	age                 int
	foreign             bool
	weinght             float64
)

func basics() {
	fmt.Printf("Basics variables \n\n")

	fmt.Printf("firstName: %s\r\n", firstName)
	fmt.Printf("lastName: %s\r\n", lastName)
	fmt.Printf("age: %d\r\n", age)
	fmt.Printf("foreign: %v\r\n", foreign)
	fmt.Printf("weinght: %f\r\n", weinght)

	fmt.Printf("\n")
	firstName = "Robert"
	lastName = "Martin"
	age = 57
	foreign = true
	weinght = 64.35

	fmt.Printf("firstName: %s\r\n", firstName)
	fmt.Printf("lastName: %s\r\n", lastName)
	fmt.Printf("age: %d\r\n", age)
	fmt.Printf("foreign: %v\r\n", foreign)
	fmt.Printf("weinght: %f\r\n", weinght)
}

func maps() {
	fmt.Printf("\nMaps \n\n")

	var address map[string]string

	fmt.Printf("address size: %d\r\n", len(address))
	fmt.Printf("address: %v\r\n", address)

	address = make(map[string]string)
	address["home"] = "Street 54, New York, number 1255"
	address["job"] = "Street 85, New York, number 19"

	fmt.Printf("address size: %d\r\n", len(address))
	fmt.Printf("address: %v\r\n", address)

	for key, value := range address {
		fmt.Printf("address: %s > %s\r\n", key, value)
	}

	delete(address, "job")
	fmt.Printf("address size: %d\r\n", len(address))
	fmt.Printf("address: %v\r\n", address)
}

func arrays() {
	fmt.Printf("\nArrays \n\n")

	var phones [3]string

	fmt.Printf("phones size: %v\r\n", len(phones))
	fmt.Printf("phones: %v\r\n", phones)

	phones[0] = "11 952368955"
	phones[1] = "11 35527748"

	fmt.Printf("phones size: %v\r\n", len(phones))
	fmt.Printf("phones: %v\r\n", phones)

	cities := [...]string{"São Paulo", "Rio Claro"}

	for i, city := range cities {
		fmt.Printf("indice [%d] is city [%s]\n", i, city)
	}
}

func slices() {
	fmt.Printf("\nSlices \n\n")

	var numbers []int
	fmt.Println(numbers, len(numbers), cap(numbers))

	numbers = make([]int, 5)
	fmt.Println(numbers, len(numbers), cap(numbers))

	cities := []string{"São Paulo", "Rio Claro", "Tokyo", "New York", "Singapura"}
	fmt.Println(cities, len(cities), cap(cities))

	cities = append(cities, "Londres")
	fmt.Println(cities, len(cities), cap(cities))

	citiesBrazil := cities[0:2]
	fmt.Println(citiesBrazil, len(citiesBrazil), cap(citiesBrazil))
}
