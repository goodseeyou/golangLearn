package main

import "fmt"

func main() {
	input := [][]int{[]int{0, 1, 0}, []int{0, 0, 1}, []int{1, 1, 1}, []int{0, 0, 0}}
	print(input)
	fmt.Println()
	gameOfLife(input)
	print(input)
}

func print(board [][]int) {
	for _, s := range board {
		fmt.Println(s)
	}
}

func gameOfLife(board [][]int) {
	for i := range board {
		for j := range board[0] {
			update(board, i, j)
		}
	}

	for i := range board {
		for j := range board[0] {
			afterUpdate(board, i, j)
		}
	}
}

func update(board [][]int, i, j int) {
	liveCount := 0

	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if (x == i && y == j) || x < 0 || x > len(board)-1 || y < 0 || y > len(board[0])-1 {
				continue
			}

			if isLive(board, x, y) {
				liveCount++
			}
		}
	}
	fmt.Println("live count", i, j, liveCount)
	originIsLive := isLive(board, i, j)
	switch {
	case liveCount < 2 && originIsLive:
		set(board, i, j, false)
	case (liveCount == 2 || liveCount == 3) && originIsLive:
		set(board, i, j, true)
	case liveCount > 3 && originIsLive:
		set(board, i, j, false)
	case !originIsLive && liveCount == 3:
		set(board, i, j, true)
	}
}

func isLive(board [][]int, i, j int) bool {
	return board[i][j] > 0
}

func set(board [][]int, i, j int, isLiveVal bool) {
	oriIsLive := isLive(board, i, j)
	if !oriIsLive && isLiveVal {
		board[i][j] = -1
	}

	if oriIsLive && !isLiveVal {
		board[i][j] = 2
	}
}

func afterUpdate(board [][]int, i, j int) {
	if board[i][j] == -1 {
		board[i][j] = 1
	}

	if board[i][j] == 2 {
		board[i][j] = 0
	}
}
