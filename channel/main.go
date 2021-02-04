package main

import (
	"fmt"
	"sync"
	"time"
)

var threadPoolCount = 10
var wg sync.WaitGroup

func r(cin <-chan int) {
	defer wg.Done()
	for item := range cin {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(item)
	}
}

func main() {
	task := make(chan int, 20)
	for i := 0; i < threadPoolCount; i++ {
		wg.Add(1)
		go r(task)
	}

	fmt.Println("start input")
	for i := 0; i < 50; i++ {
		func() { task <- i }()
	}
	fmt.Println("end input")
	close(task)
	wg.Wait()
}
