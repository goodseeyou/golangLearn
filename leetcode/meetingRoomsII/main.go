package main

import (
	"fmt"
	"sort"
)

func main() {
	input := [][]int{[]int{0, 30}, []int{5, 10}, []int{15, 20}}
	input = [][]int{[]int{13, 15}, []int{1, 13}}
	input = [][]int{[]int{3, 18}, []int{7, 16}, []int{5, 18}}
	fmt.Println(minMeetingRooms(input))
}

type OrderedStack struct {
	stack []int
}

func (this *OrderedStack) Len() int {
	return len(this.stack)
}

func (this *OrderedStack) Pop() int {
	p := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return p
}

func (this *OrderedStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *OrderedStack) Push(v int) {
	if this.Len() == 0 {
		this.stack = append(this.stack, v)
		return
	}

	for head, tail := 0, len(this.stack)-1; ; {
		if head == tail {
			insertIndex := head
			if v < this.stack[head] {
				insertIndex++
			}
			this.stack = append(this.stack[:head+1], this.stack[head:]...)
			this.stack[insertIndex] = v
			break
		}

		mid := head + (tail-head)/2
		if v > this.stack[mid] {
			tail = mid
		} else {
			head = mid + 1
		}
	}

}

func minMeetingRooms(intervals [][]int) int {
	maxRoom, curRoom := 0, 0
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	endStack := OrderedStack{stack: make([]int, 0)}
	for _, se := range intervals {
		s, e := se[0], se[1]

		for endStack.Len() > 0 && endStack.Top() <= s {
			endStack.Pop()
			curRoom--
		}

		endStack.Push(e)
		curRoom++
		if maxRoom < curRoom {
			maxRoom = curRoom
		}
	}

	return maxRoom
}
