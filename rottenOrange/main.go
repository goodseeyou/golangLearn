package main

import "fmt"

func main() {
	grid := [][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}}
	fmt.Println(orangesRotting(grid))
}

func orangesRotting(grid [][]int) int {
	rottenNodes := getRottenNodes(grid)
	rottenTime := calculateRottenTime(rottenNodes, grid)
	if !isAllRotten(grid) {
		return -1
	}

	return rottenTime
}

func getRottenNodes(grid [][]int) [][]int {
	rottenNodes := make([][]int, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				rottenNodes = append(rottenNodes, []int{i, j})
			}
		}
	}

	return rottenNodes
}

func calculateRottenTime(rottenNodes, grid [][]int) int {
	var states int
	for states = 0; len(rottenNodes) != 0; states++ {
		lenNode := len(rottenNodes)
		for i := 0; i < lenNode; i++ {
			node := rottenNodes[i]
			x, y := node[0], node[1]

			if x > 0 && grid[x-1][y] == 1 {
				rottenNodes = append(rottenNodes, []int{x - 1, y})
				grid[x-1][y] = 2
			}

			if y > 0 && grid[x][y-1] == 1 {
				rottenNodes = append(rottenNodes, []int{x, y - 1})
				grid[x][y-1] = 2
			}

			if x < len(grid)-1 && grid[x+1][y] == 1 {
				rottenNodes = append(rottenNodes, []int{x + 1, y})
				grid[x+1][y] = 2
			}

			if y < len(grid[0])-1 && grid[x][y+1] == 1 {
				rottenNodes = append(rottenNodes, []int{x, y + 1})
				grid[x][y+1] = 2
			}
		}
		rottenNodes = rottenNodes[lenNode:]
		fmt.Println(states)
	}

	if states > 0 {
		return states - 1
	}

	return 0
}

func isAllRotten(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return false
			}
		}
	}

	return true
}
