package main

type BiNode struct {
	Key  int
	Val  int
	Prev *BiNode
	Next *BiNode
}

type LRUCache struct {
	cacheMap map[int]*BiNode
	head     *BiNode // the left bound void node and won't change
	tail     *BiNode // the right bound void node and won't change
	count    int
	cap      int
}

func Constructor(capacity int) LRUCache {
	head := BiNode{0, 0, nil, nil}
	tail := BiNode{0, 0, nil, nil}
	head.Next = &tail
	tail.Prev = &head
	return LRUCache{make(map[int]*BiNode), &head, &tail, 0, capacity}
}

func (cache *LRUCache) Get(key int) int {
	if cache.cap == 0 { // will capacity be 0 ?
		return -1
	}

	if node, ok := cache.cacheMap[key]; ok {

		// pick
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		// push back
		cache.tail.Prev.Next = node
		node.Prev = cache.tail.Prev
		node.Next = cache.tail
		cache.tail.Prev = node

		return node.Val
	}

	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if cache.cap == 0 { // will capacity be 0 ?
		return
	}

	if node, ok := cache.cacheMap[key]; ok {
		// old key
		// pick
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		// push back
		cache.tail.Prev.Next = node
		node.Prev = cache.tail.Prev
		node.Next = cache.tail
		cache.tail.Prev = node

		// update value
		node.Val = value
	} else {
		// new key
		if cache.count == cache.cap {
			// delete the first node key from map
			delete(cache.cacheMap, cache.head.Next.Key)
			// pick the first node
			cache.head.Next = cache.head.Next.Next
			cache.head.Next.Prev = cache.head
			cache.count--
		}

		// add the node before tail
		newNode := BiNode{key, value, cache.tail.Prev, cache.tail}
		cache.tail.Prev.Next = &newNode
		cache.tail.Prev = &newNode
		cache.count++

		// add to map
		cache.cacheMap[key] = &newNode
	}
}

func main() {
	cache := Constructor(2 /* capacity */)

	cache.Put(1, 1)
	cache.Put(2, 2)
	println(cache.Get(1)) // returns 1
	cache.Put(3, 3)       // evicts key 2
	println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)       // evicts key 1
	println(cache.Get(1)) // returns -1 (not found)
	println(cache.Get(3)) // returns 3
	println(cache.Get(4)) // returns 4
}
