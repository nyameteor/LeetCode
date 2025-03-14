package main

import "fmt"

// ----------------------------------------------------------------------------
// Doubly linked List Implementation
// ----------------------------------------------------------------------------

// Node represents a doubly linked list node with a key-value pair
type Node struct {
	key, val   int
	prev, next *Node
}

// List represents a doubly linked list with a head and tail sentinel nodes
type List struct {
	head, tail *Node
}

func NewList() *List {
	l := &List{
		head: &Node{},
		tail: &Node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *List) insertAfterNode(prevNode, node *Node) {
	node.next = prevNode.next
	node.prev = prevNode
	prevNode.next.prev = node
	prevNode.next = node
}

// insert a node at the front of the list (after head)
func (l *List) pushFront(node *Node) {
	l.insertAfterNode(l.head, node)
}

func (l *List) removeNode(node *Node) *Node {
	node.prev.next = node.next
	node.next.prev = node.prev
	return node
}

// remove and return the last node before tail
func (l *List) popBack() *Node {
	last := l.tail.prev
	if last == l.head {
		return nil
	}
	l.removeNode(last)
	return last
}

// ----------------------------------------------------------------------------
// LRU Cache Implementation
// ----------------------------------------------------------------------------

// LRUCache represents an LRU cache with a given capacity
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	list     *List
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		list:     NewList(),
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.list.removeNode(node)
		this.list.pushFront(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.list.removeNode(node)
		this.list.pushFront(node)
		return
	}

	if len(this.cache) >= this.capacity {
		evicted := this.list.popBack()
		if evicted != nil {
			delete(this.cache, evicted.key)
		}
	}

	node := &Node{key: key, val: value}
	this.cache[key] = node
	this.list.pushFront(node)
}

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	// 1
	fmt.Println(cache.Get(1))

	cache.Put(3, 3)
	// -1
	fmt.Println(cache.Get(2))

	cache.Put(4, 4)
	// -1
	fmt.Println(cache.Get(1))
	// 3
	fmt.Println(cache.Get(3))
	// 4
	fmt.Println(cache.Get(4))
}
