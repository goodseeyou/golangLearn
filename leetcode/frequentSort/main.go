package main

import (
	"fmt"
	"sort"
	"time"
)

func frequencySort(s string) string {
	rs := []rune(s)
	cMap := make(map[rune]uint)

	for _, r := range rs {
		cMap[r]++
	}

	sort.Slice(rs, func(i, j int) bool {
		if cMap[rs[i]] == cMap[rs[j]] {
			return rs[i] > rs[j]
		}
		return cMap[rs[i]] > cMap[rs[j]]
	})

	return string(rs)
}

func frequencyBucketSort(s string) string {
	rs := []rune(s)
	cMap := make(map[rune]uint)

	maxFreq := uint(0)
	for _, r := range rs {
		cMap[r]++
		if cMap[r] > maxFreq {
			maxFreq = cMap[r]
		}
	}

	//bucket sort
	bucket := make([][]rune, maxFreq+1)
	for i := range bucket {
		bucket[i] = make([]rune, 0)
	}
	for k, v := range cMap {
		bucket[v] = append(bucket[v], k)
	}

	out := make([]rune, 0, len(s))

	for f := len(bucket) - 1; f > 0; f-- {
		rs = bucket[f]
		for _, r := range rs {
			for i := 0; i < f; i++ {
				out = append(out, r)
			}
		}
	}
	return string(out)
}

func main() {
	input := ""
	s := time.Now().UnixNano()
	fmt.Println("Result: ", frequencySort(input))
	fmt.Println("nlog(n): ", time.Now().UnixNano()-s)
	s = time.Now().UnixNano()
	fmt.Println("Result: ", frequencyBucketSort(input))
	fmt.Println("n: ", time.Now().UnixNano()-s)
}
