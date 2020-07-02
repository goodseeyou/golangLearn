package main

import "fmt"

func main() {
	numCourses := 3
	prerequisites := [][]int{[]int{1, 2}, []int{0, 1}, []int{0, 2}}
	//prerequisites = [][]int{[]int{1, 0}, []int{0, 1}}
	fmt.Println(canFinish(numCourses, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for i := range g {
		g[i] = make([]int, numCourses)
	}

	for _, p := range prerequisites {
		g[p[0]][p[1]] = 1
	}

	done := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if done[i] == 1 {
			continue
		}
		if hasCircleSeen(g, done, i) {
			return false
		}
	}

	return true
}

func hasCircleSeen(g [][]int, done []int, node int) (hasCircle bool) {
	done[node] = -1

	for i, n := range g[node] {
		if n == 0 {
			continue
		}

		if done[i] == 0 {
			if hasCircleSeen(g, done, i) {
				return true
			}
		} else if done[i] == -1 {
			return true
		}
	}
	done[node] = 1

	return false
}
