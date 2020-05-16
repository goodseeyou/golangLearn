package main

import "fmt"

func generateParenthesis(n int) []string {
	tmp := []string{""}
	for i := 0; i < n; i++ {
		m := make(map[string]byte)
		for _, s := range tmp {
			filled := fill(s)
			for k := range filled {
				m[k]++
			}
		}
		tmp = make([]string, 0)
		for k := range m {
			tmp = append(tmp, k)
		}
		fmt.Println(tmp)
	}

	return tmp
}

func fill(s string) map[string]byte {
	ss := make(map[string]byte)
	if len(s) == 0 {
		ss["()"]++
		return ss
	}
	for i := range s {
		ss[string(s[:i+1])+"()"+string(s[i+1:])]++
	}
	return ss
}

func main() {
	generateParenthesis(10)
}
