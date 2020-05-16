package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	input := [][]int{[]int{0, 1}, []int{1, 0}}
	K := 2
	fmt.Println(kClosest(input, K))
}

type Point struct {
	x, y     int
	distance float64
}

func kClosest(points [][]int, K int) [][]int {
	ps := make([]Point, len(points))
	for i, p := range points {
		ps[i] = Point{x: p[0], y: p[1], distance: distance(p)}
	}

	ans := make([][]int, K)
	bucket := []Point{}
	observ := ps
	for len(bucket) < K {
		pivot := observ[rand.Intn(len(observ))]
		less, large := []Point{}, []Point{}
		for _, p := range observ {
			if p.distance <= pivot.distance {
				less = append(less, p)
			} else {
				large = append(large, p)
			}
		}
		if len(less) > K-len(bucket) {
			observ = less
		} else if len(less) == K-len(bucket) {
			bucket = append(bucket, less...)
			break
		} else {
			bucket = append(bucket, less...)
			observ = large
		}
	}

	for i, p := range bucket {
		ans[i] = []int{p.x, p.y}
	}

	return ans
}

func distance(p []int) float64 {
	x, y := float64(p[0]), float64(p[1])
	return math.Sqrt(x*x + y*y)
}
