package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//using newsource is not thread safe for goroutine
	// i.e. two threads probably get the same random number at the same time
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		//fmt.Println(rand.Float64())
		fmt.Println(random.Float64())
	}
}
