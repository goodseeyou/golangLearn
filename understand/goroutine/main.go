package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c := make(chan string, 2)

	c <- "hello"
	close(c)

	go worker(c, wg)
	wg.Wait()
}

func worker(job chan string, wg *sync.WaitGroup) {
	for s := range job {
		fmt.Println("worker do: ", s)
	}
	wg.Done()
}
