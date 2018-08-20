package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	openTXT()

	openCSV()

	writeFile()
}

func writeFile() {

	cities := [...]string{"São Paulo", "Rio Claro", "New York", "Santos", "Tokyo"}

	citiesFile, err := os.Create("/tmp/citiesFile.txt")
	if err != nil {
		fmt.Printf("Could not create txt file err=%v\n", err)
		return
	}
	defer citiesFile.Close()

	writer := bufio.NewWriter(citiesFile)
	for _, city := range cities {
		writer.WriteString(city)
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Could not write txt file err=%v\n", err)
		return
	}

	fmt.Printf("txt file written successfully\n")

}

func openTXT() {

	citiesFile, err := os.Open("cities.txt")
	if err != nil {
		fmt.Printf("Could not open txt file err=%v\n", err)
		return
	}
	defer citiesFile.Close()

	txtScanner := bufio.NewScanner(citiesFile)
	for txtScanner.Scan() {
		line := txtScanner.Text()
		fmt.Println(line)
	}

}

func openCSV() {

	citiesFile, err := os.Open("cities.csv")
	if err != nil {
		fmt.Printf("Could not open csv file err=%v\n", err)
		return
	}
	defer citiesFile.Close()

	reader := csv.NewReader(citiesFile)
	content, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Could not read csv err=%v\n", err)
		return
	}

	for i, line := range content {
		fmt.Printf("index [%d] content [%s]\n", i, line)
	}

}
