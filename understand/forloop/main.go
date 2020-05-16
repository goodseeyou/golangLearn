package main

import "fmt"

func main() {
	var i int
	//s := [][]int{[]int{1}, []int{1}, []int{1}, []int{1}, []int{1}}
	s := [][]int{[]int{1}, []int{1}}
	for i = 0; len(s) != 0; i++ {
		// add an element
		s = append(s, []int{2})
		// remove two elements
		s = s[2:]
		fmt.Println(i)
	}
	fmt.Println("finally: ", i)
}
