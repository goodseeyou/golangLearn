package main

import "fmt"

func main() {
	input := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println(spiralOrder(input))
}

var directions = [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}

const lenDirections = 4

var directionsIndexMap = map[[2]int]int{[2]int{0, 1}: 0, [2]int{1, 0}: 1, [2]int{0, -1}: 2, [2]int{-1, 0}: 3}

func spiralOrder(matrix [][]int) []int {
	var M = len(matrix)
	if M == 0 {
		return []int{}
	}
	var N = len(matrix[0])
	if N == 0 {
		return []int{}
	}

	output := make([]int, 0, M*N)
	doneMap := make(map[[2]int]bool)
	direct := directions[0]
	isContinue := true
	for i, j := 0, 0; isContinue; {
		output = append(output, matrix[i][j])
		doneMap[[2]int{i, j}] = true
		direct, isContinue = getNextPoint(i, j, M, N, direct, doneMap)
		i += direct[0]
		j += direct[1]
	}
	return output
}

func getNextPoint(i, j, M, N int, lastDirect [2]int, doneMap map[[2]int]bool) ([2]int, bool) {
	lastDirectionsIndex := directionsIndexMap[lastDirect]

	nextI, nextJ := i+lastDirect[0], j+lastDirect[1]
	if isInvalid(nextI, nextJ, M, N, doneMap) {
		nextDirect := directions[(lastDirectionsIndex+1)%lenDirections]
		nextI, nextJ = i+nextDirect[0], j+nextDirect[1]
		if isInvalid(nextI, nextJ, M, N, doneMap) {
			return nextDirect, false
		}
		return nextDirect, true
	}

	return lastDirect, true
}

func isInvalid(i, j, M, N int, doneMap map[[2]int]bool) bool {
	//fmt.Println("doneMap:", i, j, doneMap[[2]int{i, j}])
	return doneMap[[2]int{i, j}] || i < 0 || i > M-1 || j < 0 || j > N-1
}
