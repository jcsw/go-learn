package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("value in 'i' is %d\n", i)
	}

	x := 10
	for x > 0 {
		fmt.Printf("value in 'x' is %d\n", x)
		x--
	}

	text := "Learning Go"

	for i, letter := range text {
		fmt.Printf("Index %d letter %q\n", i, letter)
	}
}
