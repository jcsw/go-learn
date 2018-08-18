package main

import (
	"errors"
	"fmt"
)

var errMyError = errors.New("My error it's here")

func main() {

	err := anything()
	if err != nil {
		fmt.Printf("Error message is '%s'\n", err.Error())
	}

	if err := anythingTwo(); err != nil {
		if err == errMyError {
			fmt.Printf("Error message is '%s'\n", err.Error())
		}
	}

}

func anything() error {
	return errors.New("Have something wrong here")
}

func anythingTwo() error {
	return errMyError
}
