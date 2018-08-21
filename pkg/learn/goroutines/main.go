package main

import (
	"fmt"
	"sync"
	"time"
)

var orchestrator sync.WaitGroup

func main() {
	startTime := time.Now()

	orchestrator.Add(4)

	allPairNumbers := [4001]int{}

	go fillPairNumbers(&allPairNumbers, 0, 1000)
	go fillPairNumbers(&allPairNumbers, 1001, 2000)
	go fillPairNumbers(&allPairNumbers, 2001, 3000)
	go fillPairNumbers(&allPairNumbers, 3001, 4000)

	orchestrator.Wait()

	printPairNumbers(allPairNumbers)

	fmt.Printf("elipsedTime: %v\n", time.Since(startTime))
}

func fillPairNumbers(pairNumbers *[4001]int, start, end int) {
	for i := start; i <= end; i++ {
		time.Sleep(1000)
		if i%2 == 0 {
			pairNumbers[i] = i
		}
	}

	orchestrator.Done()
}

func printPairNumbers(pairNumbers [4001]int) {
	for _, n := range pairNumbers {
		if n != 0 {
			fmt.Printf("[printPairNumbers] number [%d] is pair\n", n)
		}
	}
}
