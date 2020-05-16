package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(nums))
}

func singleNonDuplicateS(nums []int) int {
	tail := len(nums) - 1
	mid := tail / 2

	if tail == 0 {
		return nums[tail]
	}

	// number of left and right part is odd
	if mid%2 == 1 {
		if nums[mid] == nums[mid+1] {
			return singleNonDuplicate(nums[:mid])
		} else {
			return singleNonDuplicate(nums[mid+1:])
		}
	} else {
		if nums[mid] == nums[mid-1] {
			return singleNonDuplicate(nums[:mid+1])
		} else {
			return singleNonDuplicate(nums[mid:])
		}
	}
}

func singleNonDuplicate(nums []int) int {
	for head, tail := 0, len(nums)-1; ; {
		mid := head + (tail-head)/2
		fmt.Println(head, mid, tail)
		if tail == head {
			return nums[tail]
		}

		if mid%2 == 1 {
			if nums[mid] == nums[mid+1] {
				tail = mid - 1
			} else {
				head = mid + 1
			}
		} else {
			if nums[mid] == nums[mid-1] {
				tail = mid
			} else {
				head = mid
			}
		}
	}
}
