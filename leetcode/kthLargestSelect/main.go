package main

import "math/rand"

func main() {

}
func findKthLargestRecur(nums []int, k int) int {
	selectIndex := rand.Intn(len(nums))
	selectValue := nums[selectIndex]
	nums[selectIndex], nums[len(nums)-1] = nums[len(nums)-1], nums[selectIndex]
	nums = nums[:len(nums)-1]

	leIndexEnd := 0
	for i, n := range nums {
		if n <= selectValue {
			nums[i], nums[leIndexEnd] = nums[leIndexEnd], nums[i]
			leIndexEnd++
		}
	}

	le := nums[:leIndexEnd]
	bt := nums[leIndexEnd:]

	if len(bt) >= k {
		return findKthLargest(bt, k)
	}

	if k-len(bt) == 1 {
		return selectValue
	}

	return findKthLargest(le, k-len(bt)-1)
}

func findKthLargest(nums []int, k int) int {
	for {
		selectIndex := rand.Intn(len(nums))
		selectValue := nums[selectIndex]
		nums[selectIndex], nums[len(nums)-1] = nums[len(nums)-1], nums[selectIndex]
		nums = nums[:len(nums)-1]

		leIndexEnd := 0
		for i, n := range nums {
			if n <= selectValue {
				nums[i], nums[leIndexEnd] = nums[leIndexEnd], nums[i]
				leIndexEnd++
			}
		}

		bt := nums[leIndexEnd:]

		if len(bt) >= k {
			nums = nums[leIndexEnd:]
			continue
		}

		if k-len(bt) == 1 {
			return selectValue
		}

		nums = nums[:leIndexEnd]
		k = k - len(bt) - 1
	}
}
