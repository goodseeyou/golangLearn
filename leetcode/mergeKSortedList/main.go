package main

import (
	"container/heap"
	"fmt"
)

func main() {
	klist := newList()
	printList(mergeKLists(klist))
	klist = newList()
	printList(mergeKListsOri(klist))
}

func newList() []*ListNode {
	//[[1,4,5],[1,3,4],[2,6]]
	klist := make([]*ListNode, 0)
	klist = append(klist, parse([]int{1, 4, 5}))
	klist = append(klist, parse([]int{1, 3, 4}))
	klist = append(klist, parse([]int{2, 6}))
	return klist
}

func printList(l *ListNode) {
	n := l
	for n != nil {
		fmt.Printf("%d, ", n.Val)
		n = n.Next
	}
	fmt.Println()
}

func parse(nums []int) *ListNode {
	head := &ListNode{}
	hold := head
	for _, n := range nums {
		hold.Next = &ListNode{Val: n}
		hold = hold.Next
	}

	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKListsOri(lists []*ListNode) *ListNode {
	head := &ListNode{}
	hold := head
	for {
		minVal := 1<<31 - 1
		var minList *ListNode = nil
		var index int
		for i, l := range lists {
			if l == nil {
				continue
			}

			if l.Val < minVal {
				minVal = l.Val
				minList = l
				index = i
			}
		}
		if minList == nil {
			break
		}
		lists[index] = minList.Next
		hold.Next = minList
		hold = hold.Next
	}

	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, list := range lists {
		if list == nil {
			continue
		}

		item := Item{value: list,
			priority: -list.Val}
		heap.Push(&pq, &item)
	}

	head := &ListNode{}
	hold := head

	for pq.Len() != 0 {
		minItem := heap.Pop(&pq).(*Item)

		hold.Next = minItem.value
		hold = hold.Next
		if minItem.value.Next != nil {
			minItem.value = minItem.value.Next
			minItem.priority = -minItem.value.Val
			heap.Push(&pq, minItem)
		}

	}

	return head.Next
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    *ListNode // The value of the item; arbitrary.
	priority int       // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value *ListNode, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
