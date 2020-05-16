package main

import "fmt"

func main() {
	head := &Trie{index: make(map[byte]*Trie),
		leave: make(map[byte]byte)}
	head.Insert("apple")
	head.Search("apple")
	head.Search("app")
	fmt.Println(head.StartsWith("app"))
	head.Insert("app")
	head.Search("app")

}

type Trie struct {
	leave map[byte]byte
	index map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{index: make(map[byte]*Trie),
		leave: make(map[byte]byte)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	if len(word) == 1 {
		this.leave[word[0]] = byte(1)
		return
	}

	if _, ok := this.index[word[0]]; !ok {
		this.index[word[0]] = &Trie{index: make(map[byte]*Trie),
			leave: make(map[byte]byte)}
	}

	this.index[word[0]].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 1 {
		_, ok := this.leave[word[0]]
		return ok
	}

	sub, ok := this.index[word[0]]
	if !ok {
		return false
	}

	return sub.Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {

	if len(prefix) == 0 {
		return false
	}

	if len(prefix) == 1 {
		_, leaveOk := this.leave[prefix[0]]
		_, indexOK := this.index[prefix[0]]
		return leaveOk || indexOK
	}

	sub, ok := this.index[prefix[0]]
	if !ok {
		return false
	}
	return sub.StartsWith(prefix[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
