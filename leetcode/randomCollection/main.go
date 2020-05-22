package main

import (
	"math/rand"
)

func main() {}

type RandomizedCollection struct {
	vMap  map[int]map[int]byte
	vList []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{vMap: make(map[int]map[int]byte), vList: make([]int, 0)}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	_, ok := this.vMap[val]
	if !ok {
		this.vMap[val] = make(map[int]byte)
	}
	this.vList = append(this.vList, val)
	this.vMap[val][len(this.vList)-1]++
	//fmt.Println("insert val:", val, "remain: ", this.vList)
	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	iMap, ok := this.vMap[val]
	if ok {
		index := -1
		for k := range iMap {
			index = k
			break
		}

		/*fmt.Print("val:", val, "| ")
		  for k := range iMap {
		      fmt.Print(k, " ")
		  }
		  fmt.Println("| index:", index)*/

		delete(this.vMap[val], index)
		this.vMap[this.vList[len(this.vList)-1]][index]++
		delete(this.vMap[this.vList[len(this.vList)-1]], len(this.vList)-1)

		//fmt.Println("map val: ", val, "len map[val]", this.vMap[val][index])
		if len(this.vMap[val]) == 0 {
			delete(this.vMap, val)
		}

		this.vList[index], this.vList[len(this.vList)-1] = this.vList[len(this.vList)-1], this.vList[index]
		this.vList = this.vList[:len(this.vList)-1]

		//fmt.Println("del val:", val, " remain: ", this.vList)
		return true
	}
	//fmt.Println("del val:", val, " remain: ", this.vList)
	return false
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	if len(this.vList) <= 0 {
		return -1
	}
	return this.vList[rand.Intn(len(this.vList))]

}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
