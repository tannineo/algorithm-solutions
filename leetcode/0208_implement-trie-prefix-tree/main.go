package main

type Trie struct {
	IsWord  bool
	ArrNext []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{false, make([]*Trie, 26)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curNode := this
	for _, v := range word {
		if curNode.ArrNext[v-'a'] == nil {
			// append trie
			curNode.ArrNext[v-'a'] = &Trie{false, make([]*Trie, 26)}
		}
		curNode = curNode.ArrNext[v-'a']
	}
	curNode.IsWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curNode := this
	for _, v := range word {
		if curNode.ArrNext[v-'a'] == nil {
			return false
		}
		curNode = curNode.ArrNext[v-'a']
	}
	return curNode.IsWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	curNode := this
	for _, v := range prefix {
		if curNode.ArrNext[v-'a'] == nil {
			return false
		}
		curNode = curNode.ArrNext[v-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()

	trie.Insert("apple")
	println(trie.Search("apple"))   // returns true
	println(trie.Search("app"))     // returns false
	println(trie.StartsWith("app")) // returns true
	trie.Insert("app")
	println(trie.Search("app")) // returns true
}
