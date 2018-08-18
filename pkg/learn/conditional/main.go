package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	randSource := rand.NewSource(time.Now().UnixNano())
	number := rand.New(randSource).Intn(100)

	// > >= < <= == != || &&
	if number > 50 {
		fmt.Printf("number %d is greater than 50\n", number)
	}

	if number >= 50 {
		fmt.Printf("number %d is greater or equals than 50\n", number)
	}

	if number < 50 {
		fmt.Printf("number %d is less than 50\n", number)
	}

	if number <= 50 {
		fmt.Printf("number %d is less or equals than 50\n", number)
	}

	if number == 50 {
		fmt.Printf("number %d is equals than 50\n", number)
	}

	if number != 50 {
		fmt.Printf("number %d is different than 50\n", number)
	}

	numberSwitch := rand.New(randSource).Intn(10)
	switch numberSwitch {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	case 7:
		fmt.Println("Seven")
	case 8:
		fmt.Println("Eight")
	case 9:
		fmt.Println("Nine")
	case 10:
		fmt.Println("Ten")
	}

	env := runtime.GOOS
	switch env {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Printf("Yes %s is Unix family\n", env)
	default:
		fmt.Printf("No %s is not Unix family\n", env)
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Make up at 10am")
	default:
		fmt.Println("Make up at 7am")
	}

	switch {
	case true:
		fmt.Println("Yes conditional at 'case'")
	}
}
