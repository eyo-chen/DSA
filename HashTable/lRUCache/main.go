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
