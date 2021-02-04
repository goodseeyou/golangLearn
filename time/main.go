package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	time.Sleep(500 * time.Millisecond)
	b := time.Now()

	fmt.Println(a.UnixNano() < b.UnixNano(), a.UnixNano(), b.UnixNano())
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())
}
