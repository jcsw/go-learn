package main

import "fmt"

type veicle struct {
	name  string
	model string
	year  int
}

func main() {

	myCar := veicle{name: "Celta", model: "Hatch"}
	myCar.setYear(2011)

	fmt.Printf("My car is %+v\n", myCar)
}

func (v *veicle) setYear(year int) {
	v.year = year
}
