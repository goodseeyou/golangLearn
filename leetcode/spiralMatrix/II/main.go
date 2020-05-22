package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

var directions = [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}

const lenDirections = 4

var directionsIndexMap = map[[2]int]int{[2]int{0, 1}: 0, [2]int{1, 0}: 1, [2]int{0, -1}: 2, [2]int{-1, 0}: 3}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	x, y := 0, 0
	direct := directions[0]
	doneMap := make(map[[2]int]bool)
	for i := 1; i <= n*n; i++ {
		matrix[x][y] = i
		doneMap[[2]int{x, y}] = true
		direct = getNextDirect(x, y, n, n, direct, doneMap)
		x += direct[0]
		y += direct[1]
	}
	return matrix
}

func getNextDirect(i, j, M, N int, lastDirect [2]int, doneMap map[[2]int]bool) [2]int {
	lastDirectionsIndex := directionsIndexMap[lastDirect]

	nextI, nextJ := i+lastDirect[0], j+lastDirect[1]
	if isInvalid(nextI, nextJ, M, N, doneMap) {
		nextDirect := directions[(lastDirectionsIndex+1)%lenDirections]
		nextI, nextJ = i+nextDirect[0], j+nextDirect[1]
		if isInvalid(nextI, nextJ, M, N, doneMap) {
			return nextDirect
		}
		return nextDirect
	}

	return lastDirect
}

func isInvalid(i, j, M, N int, doneMap map[[2]int]bool) bool {
	return doneMap[[2]int{i, j}] || i < 0 || i > M-1 || j < 0 || j > N-1
}
