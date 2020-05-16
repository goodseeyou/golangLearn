package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5}
	head := gen(input)
	print(head)
	head = oddEvenList(head)
	print(head)
}

func print(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, ",")
		node = node.Next
	}
	fmt.Println()
}

func gen(nums []int) *ListNode {
	head := &ListNode{}
	hold := head
	for _, n := range nums {
		node := &ListNode{Val: n}
		hold.Next = node
		hold = hold.Next
	}

	return head.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	hold := head

	var odd *ListNode = hold
	hold = hold.Next
	if hold == nil {
		return odd
	}

	var even *ListNode = hold
	evenHead := hold
	hold = hold.Next

	for i := 1; hold != nil; i++ {
		if i%2 == 1 {
			odd.Next = hold
			odd = odd.Next
		} else {
			even.Next = hold
			even = even.Next
		}

		hold = hold.Next
	}
	even.Next = nil
	odd.Next = evenHead

	return head
}
