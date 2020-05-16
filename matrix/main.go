package main

import "fmt"

//var I = [][]int{[]int{1, 0, 1, 0, 0}, []int{1, 0, 1, 1, 0}, []int{1, 1, 1, 1, 1}, []int{1, 0, 0, 1, 0}}
//var I = [][]int{[]int{1,0}, []int{1,0}}
var I = [][]int{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}, []int{1, 1, 1, 1}}

func main() {
	input := make([][]byte, len(I))
	for i := range I {
		input[i] = make([]byte, len(I[0]))
		for j := range I[0] {
			input[i][j] = byte(I[i][j])
		}
	}
	fmt.Println(I)
	fmt.Println(input)

	fmt.Println(maximalSquare(input))
}

func maximalSquare(matrix [][]byte) int {
	record := make([][]int, len(matrix))
	for i := range matrix {
		record[i] = make([]int, len(matrix[0]))
		for j := range matrix[0] {
			record[i][j] = -1
		}
	}

	ans := 0
	for i := range matrix {
		for j := range matrix[0] {
			v := r(matrix, i, j, record)
			if v > ans {
				ans = v
			}
		}
	}
	fmt.Println("record:", record)

	return ans * ans
}

func r(m [][]byte, i, j int, record [][]int) int {
	if record[i][j] >= 0 {
		return record[i][j]
	}

	//if string(m[i][j]) != "1" {
	if m[i][j] != byte(1) {
		record[i][j] = 0
		return 0
	}

	if i < len(m)-1 {
		record[i+1][j] = r(m, i+1, j, record)
	} else {
		record[i][j] = 1
		return 1
	}

	if j < len(m[0])-1 {
		record[i][j+1] = r(m, i, j+1, record)
	} else {
		record[i][j] = 1
		return 1
	}

	record[i+1][j+1] = r(m, i+1, j+1, record)
	minv := min(record[i+1][j], record[i][j+1], record[i+1][j+1])
	record[i][j] = minv + 1
	return record[i][j]
}

func max(nums ...int) int {
	maxv := -1 << 63

	for _, v := range nums {
		if v > maxv {
			maxv = v
		}
	}

	return maxv
}

func min(nums ...int) int {
	minv := 1<<63 - 1

	for _, v := range nums {
		if v < minv {
			minv = v
		}
	}

	return minv
}
