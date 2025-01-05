package main

type doubleLinkedList struct {
	val  int
	key  int // we need the key to remove the node from the hash table because delete() function of map needs the key
	prev *doubleLinkedList
	next *doubleLinkedList
}

type LRUCache struct {
	// the value has to be the node of double linked list
	// because for both get() and put() function, we need to update the node's position once the node is accessed
	hashTable map[int]*doubleLinkedList
	cap       int

	// for both head and tail, it's a dummy node
	// it's convenient for the insert and remove operation
	head *doubleLinkedList
	tail *doubleLinkedList
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		hashTable: make(map[int]*doubleLinkedList),
		cap:       capacity,
		head:      &doubleLinkedList{},
		tail:      &doubleLinkedList{},
	}

	// connect head and tail
	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru
}

func (l *LRUCache) Get(key int) int {
	// if the key is in the hash table, update the node's position to the head
	if v, ok := l.hashTable[key]; ok {
		l.removeNode(v)
		l.addToHead(v)
		return v.val
	}

	// if the key is not in the hash table, return -1
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	// if the key is in the hash table
	// update the node's value
	// update the node's position to the head
	if v, ok := l.hashTable[key]; ok {
		v.val = value
		l.removeNode(v)
		l.addToHead(v)
		return
	}

	// if the key is not in the hash table
	// and the hash table is full, remove the tail node
	if len(l.hashTable) == l.cap {
		tail := l.popTail()
		delete(l.hashTable, tail.key)
	}

	// add the new node to the head
	node := &doubleLinkedList{key: key, val: value}
	l.hashTable[key] = node
	l.addToHead(node)
}

func (l *LRUCache) removeNode(node *doubleLinkedList) {
	// reference to the previous and next node
	prev := node.prev
	next := node.next

	// update the previous and next node
	prev.next = next
	next.prev = prev
}

func (l *LRUCache) addToHead(node *doubleLinkedList) {
	// make current node connect to the head
	node.prev = l.head

	// make current node connect to the the old first node
	node.next = l.head.next

	// make the head connect to the current node
	l.head.next = node

	// make the old first node connect to the current node
	node.next.prev = node
}

func (l *LRUCache) popTail() *doubleLinkedList {
	tail := l.tail.prev
	l.removeNode(tail)
	return tail
}

// updated at 2025-01-05
type doubleLinkedList1 struct {
	val  int
	key  int
	prev *doubleLinkedList1
	next *doubleLinkedList1
}

type LRUCache1 struct {
	hashTable map[int]*doubleLinkedList1
	cap       int
	head      *doubleLinkedList1 // least used
	tail      *doubleLinkedList1 // most used
}

func Constructor1(capacity int) LRUCache1 {
	// Initialize the head and tail
	head := &doubleLinkedList1{}
	tail := &doubleLinkedList1{}

	// Connect the head and tail
	head.next = tail
	tail.prev = head

	return LRUCache1{
		hashTable: make(map[int]*doubleLinkedList1, capacity),
		cap:       capacity,
		head:      head,
		tail:      tail,
	}
}

func (l *LRUCache1) Get(key int) int {
	// Get the node from the hash table
	// If the node is not found, return -1
	node, exist := l.hashTable[key]
	if !exist {
		return -1
	}

	// Update the node to the tail
	l.updateToTail(node)

	// Return the value of the node
	return node.val
}

func (l *LRUCache1) Put(key int, value int) {
	// If the node is found, update the value and move it to the tail
	if node, exist := l.hashTable[key]; exist {
		// update the value
		node.val = value

		// update the node to the tail
		l.updateToTail(node)
		return
	}

	// If the hash table is full, remove the least used node
	if len(l.hashTable) == l.cap {
		// get the least used node (head is the least used node)
		leastUsedNode := l.head.next

		// remove the least used node
		l.removeNode(leastUsedNode)

		// delete the least used node from the hash table
		delete(l.hashTable, leastUsedNode.key)
	}

	// add the new node to the hash table
	node := &doubleLinkedList1{val: value, key: key}
	l.hashTable[key] = node

	// add the new node to the tail
	l.addToTail(node)
}

func (l *LRUCache1) updateToTail(node *doubleLinkedList1) {
	// remove the node from the list
	l.removeNode(node)

	// add the node to the tail
	l.addToTail(node)
}

func (l *LRUCache1) addToTail(node *doubleLinkedList1) {
	// connect the node to the tail
	node.next = l.tail
	node.prev = l.tail.prev

	// connect the tail to the node
	l.tail.prev.next = node
	l.tail.prev = node
}

func (l *LRUCache1) removeNode(node *doubleLinkedList1) {
	// get the next and previous node
	next := node.next
	prev := node.prev

	// update the next and previous node
	prev.next = next
	next.prev = prev

	// disconnect the node
	node.next = nil
	node.prev = nil
}
