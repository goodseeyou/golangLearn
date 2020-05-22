package main

import "fmt"

func main() {
	hc := Constructor()
	hc.Hit(1)
	hc.Hit(2)
	hc.Hit(3)
	fmt.Println(hc.GetHits(4))
	hc.Hit(300)
	fmt.Println(hc.GetHits(300))
	fmt.Println(hc.GetHits(301))

}

const duration = 60 * 5

type HitCounter struct {
	head, tail *LinkedList
	nodeMap    map[int]*LinkedList
	count      int
}

type LinkedList struct {
	left, right *LinkedList
	count       int
	timestamp   int
}

func (this *LinkedList) InsertL(node *LinkedList) {
	oldLeft := this.left
	this.left = node
	oldLeft.right = node
	node.right = this
	node.left = oldLeft
}

func (this *LinkedList) removeSelf() {
	this.left.right = this.right
	this.right.left = this.left
	this.right = nil
	this.left = nil
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	nodeMap := make(map[int]*LinkedList)
	head := &LinkedList{}
	tail := &LinkedList{}
	head.left = head
	head.right = tail
	tail.left = head
	tail.right = tail
	return HitCounter{head: head, tail: tail, nodeMap: nodeMap}

}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	node, ok := this.nodeMap[timestamp]
	if !ok {
		node = &LinkedList{timestamp: timestamp}
		this.tail.InsertL(node)
	}
	node.count++
	this.count++
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	node := this.head.right
	for this.count > 0 && node.timestamp <= timestamp-duration {
		this.count -= node.count
		nextNode := node.right
		node.removeSelf()
		delete(this.nodeMap, node.timestamp)
		node = nextNode
	}

	return this.count
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
