package main

import "fmt"

func main() {
    l := Constructor(2)
    l.Put(1,1)
    l.P()
    l.Put(2,2)
    l.P()
    fmt.Println(l.Get(1))
    l.P()
    l.Put(3,3)
    fmt.Println(l.Get(2))
    l.P()
}

type LRUCache struct {
    mapToLinkedList map[int]*LinkedList
    m map[int]int
    capacity int
    head, tail *LinkedList
}

func (this *LRUCache) P() {
    this.head.Right.Println()
}

func Constructor(capacity int) LRUCache {
    lruCache := LRUCache{
        capacity:capacity,
        mapToLinkedList: make(map[int]*LinkedList, capacity),
        m: make(map[int]int, capacity),
        head: &LinkedList{},
        tail: &LinkedList{},
    }
    lruCache.Init()

    return lruCache
}

func (this *LRUCache) Init() {
    this.head.Left = this.head
    this.head.Right = this.tail
    this.tail.Left = this.head
    this.tail.Right = this.tail
}


func (this *LRUCache) Get(key int) int {
    if v, ok := this.m[key]; ok {
        this.use(key)
        return v
    }

    return -1
}

func (this *LRUCache) Put(key int, value int)  {
    _, ok := this.m[key]
    if !ok && this.capacity <= len(this.m) {
        this.removeLeastUsedKey()
    }
    this.m[key] = value
    this.addNewKey(key)
}

func (this *LRUCache) removeLeastUsedKey(){
    leastUsedItem := this.tail.Left
    defer delete(this.mapToLinkedList, leastUsedItem.Key)
    defer delete(this.m, leastUsedItem.Key)
    this.tail.Left = leastUsedItem.Left
    this.tail.Left.Right = this.tail
}

func (this *LRUCache) addNewKey(key int) {
    lastUsedItem := this.head.Right
    newList := &LinkedList{Key:key}
    this.mapToLinkedList[key] = newList
    linkTwoLinkedList(this.head, newList)
    linkTwoLinkedList(newList, lastUsedItem)
}

func (this *LRUCache) use(key int) {
    lastUsedItem := this.head.Right
    usedItem := this.mapToLinkedList[key]

    if lastUsedItem.Key == usedItem.Key {
        return
    }

    // strict process order
    linkTwoLinkedList(usedItem.Left, usedItem.Right)
    linkTwoLinkedList(this.head, usedItem)
    linkTwoLinkedList(usedItem, lastUsedItem)
}

type LinkedList struct {
    Right, Left *LinkedList
    Key int
}

func linkTwoLinkedList(l1, l2 *LinkedList){
    l1.Right = l2
    l2.Left = l1
}

func (this *LinkedList) Println(){
    for hold := this ; hold.Right!= hold ; hold = hold.Right {
        fmt.Print(" ", hold.Key, " ")
    }
    fmt.Println()
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
