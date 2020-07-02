package main

import "fmt"

func main() {
	T := []int{10, -2, 2, 5, 2, 5, 9, 8, 19}
	fmt.Println(minWinter(T))
}

func minWinter(T []int) int {
	minIs := make([]int, len(T))
	minV := T[len(T)-1]
	for i := len(T) - 1; i > 0; i-- {
		if T[i] <= minV {
			minV = T[i]
		}
		minIs[i-1] = minV
	}
	fmt.Println(minIs)
	maxIs := make([]int, len(T))
	maxV := T[0]
	for i := 0; i < len(T); i++ {
		if T[i] > maxV {
			maxV = T[i]
		}
		maxIs[i] = maxV

		if maxV < minIs[i] {
			return i + 1
		}
	}
	fmt.Println(maxIs)
	return -1
}
