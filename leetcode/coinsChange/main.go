package main

import "fmt"

func main() {
	coins := []int{86, 419, 83, 408}
	amount := 6249
	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {

	table := make([]int, amount+1)
	for i := range table {
		table[i] = 1<<31 - 1
	}
	table[0] = 0

	for _, c := range coins {
		for i := c; i <= amount; i++ {
			table[i] = Min(table[i], table[i-c]+1)
		}
	}

	if table[amount] == 1<<31-1 {
		return -1
	}

	return table[amount]
}

func Min(nums ...int) int {
	min := nums[0]

	for _, n := range nums {
		if min > n {
			min = n
		}
	}

	return min
}
