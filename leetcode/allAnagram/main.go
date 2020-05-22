package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	pMap := make(map[byte]int)

	for i := range p {
		pMap[p[i]]++
	}

	output := make([]int, 0)
	for i, j := 0, 0; i <= len(s)-len(p) && j < len(s); {
		pVal, ok := pMap[s[j]]
		if !ok {
			j++
			for ; i < j; i++ {
				_, ok := pMap[s[i]]
				if ok {
					pMap[s[i]]++
				}
			}
			continue
		}

		if pVal > 0 {
			pMap[s[j]]--
			j++
			if (j - i) == len(p) {
				output = append(output, i)
				pMap[s[i]]++
				i++
			}
			continue
		}

		if i == j {
			j++
		}
		pMap[s[i]]++
		i++

	}

	return output
}
