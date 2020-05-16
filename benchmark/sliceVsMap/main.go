package main

import (
	"fmt"
	"math/rand"
	"time"
)

const RunTimes = 10000
const RandIntUpperBound = 511

func main() {
	beginTime := time.Now()
	runTimes(sliceCreateAndSum(), 100)
	fmt.Println("slice time: ", time.Now().Sub(beginTime))

	beginTime = time.Now()
	runTimes(mapCreateAndSum(), 100)
	fmt.Println("map time: ", time.Now().Sub(beginTime))
}

func sliceCreateAndSum() func() {
	return func() {
		s := make([]int, 0)
		for i := 0; i < RunTimes; i++ {
			s = append(s, getRandInt())
		}
		sum := 0
		for _, v := range s {
			sum += v
		}
	}
}

func mapCreateAndSum() func() {
	return func() {
		m := make(map[int]int)

		for i := 0; i < RunTimes; i++ {
			m[i] = getRandInt()
		}
		sum := 0
		for _, v := range m {
			sum += v
		}
	}
}

func getRandInt() int {
	return rand.Int() % RandIntUpperBound
}

func runTimes(f func(), times int) {
	for i := 0; i < times; i++ {
		f()
	}
}
