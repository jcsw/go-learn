package main

import (
	"fmt"
	"sync"
	"time"
)

var workers = 4
var tasks = make(chan int, workers)
var wg = sync.WaitGroup{}
var numbers = 15000

func main() {
	startTime := time.Now()

	wg.Add(workers)
	for worker := 1; worker <= workers; worker++ {
		go process(worker)
	}

	for {
		tasks <- numbers
		numbers--
		if numbers == 0 {
			break
		}
	}
	close(tasks)
	wg.Wait()

	fmt.Printf("elipsed time: %v\n", time.Since(startTime))
}

func process(worker int) {
	defer wg.Done()
	fmt.Printf("init worker:%d\n", worker)
	count := 0
	for {
		task, more := <-tasks
		if more {
			processNumber(task)
			count++
		} else {
			fmt.Printf("received all tasks, worker:%d count:%d\n", worker, count)
			break
		}
	}
}

func processNumber(number int) {
	if number%2 == 0 {
		time.Sleep(time.Duration(int64(number)))
	}
}
