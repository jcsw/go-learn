package main

import "fmt"

type bird interface {
	fly() (string, error)
}

type eagle struct {
	name string
}

func (e eagle) fly() (string, error) {
	return fmt.Sprintf("bird [%s] it's flying", e.name), nil
}

func main() {
	louise := eagle{name: "louise"}
	lowFlight(louise)
}

func lowFlight(b bird) {
	fly, _ := b.fly()
	fmt.Printf(fly + ", low flight\n")
}
