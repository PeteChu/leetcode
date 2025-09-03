package lrucache

type Node struct {
	prev  *Node
	key   int
	value int
	next  *Node
}

func (n Node) Value() int {
	return n.value
}

type LRUCache struct {
	cache    map[int]*Node
	capacity int
	size     int
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*Node),
		capacity: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.cache[key]
	if !ok {
		return -1
	}

	if node != l.head {
		l.moveToHead(node)
	}
	return node.value
}

func (l *LRUCache) Put(key int, value int) {
	// if key existed update it
	if node, ok := l.cache[key]; ok {
		node.value = value
		l.moveToHead(node)
		return
	}

	// check if new key exceed capacity
	l.size++
	if l.size > l.capacity {
		key := l.tail.key
		if l.tail.prev != nil {
			l.tail.prev.next = nil
		}
		l.tail = l.tail.prev
		delete(l.cache, key)
		l.size--
	}

	// add new value to head
	node := &Node{
		key:   key,
		value: value,
		next:  l.head,
	}

	// update head and tail
	switch l.size {
	case 1:
		l.head = node
		l.tail = node
	case 2:
		if l.tail != nil {
			l.tail.prev = node
		}
		node.next = l.tail
		l.head = node
	default:
		l.head.prev = node
		node.next = l.head
		l.head = node
	}
	l.cache[key] = node
}

func (l *LRUCache) moveToHead(node *Node) {
	if l.head == node || l.head == nil {
		return
	}

	if node == l.tail {
		// new tail
		node.prev.next = nil
		l.tail = node.prev
	}

	node.prev.next = node.next

	if node.next != nil {
		node.next.prev = node.prev
	}

	// move node to head
	node.next = l.head
	node.prev = nil
	l.head.prev = node
	l.head = node
}

func (l *LRUCache) Capacity() int {
	return l.capacity
}

func (l *LRUCache) Size() int {
	return l.size
}

func (l *LRUCache) Head() Node {
	return *l.head
}

func (l *LRUCache) Values() []int {
	values := []int{}

	node := l.head
	for node != nil {
		values = append(values, node.value)
		node = node.next
	}

	return values
}
