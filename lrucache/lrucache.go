package lrucache

type Node[T comparable] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type LRUcache[K comparable, V comparable] struct {
	capacity     int
	length       int
	head         *Node[V]
	tail         *Node[V]
	cache        map[K]*Node[V]
	reverseCache map[*Node[V]]K
}

func New[K comparable, V comparable](cap int) *LRUcache[K, V] {
	return &LRUcache[K, V]{
		capacity:     cap,
		length:       0,
		head:         nil,
		tail:         nil,
		cache:        make(map[K]*Node[V]),
		reverseCache: make(map[*Node[V]]K),
	}
}

func (c *LRUcache[K, V]) Len() int {
	return c.length
}

func (c *LRUcache[K, V]) Update(key K, value V) {
	node := c.cache[key]
	if node == nil {
		node = &Node[V]{value: value}
		c.length++
		c.prepend(node)
		c.trimcache()
		c.cache[key] = node
		c.reverseCache[node] = key
	} else {
		node.value = value
		c.detach(node)
		c.prepend(node)
	}
}

func (c *LRUcache[K, V]) Get(key K) (val V, ok bool) {
	node := c.cache[key]
	if node == nil {
		return val, false
	}

	c.detach(node)
	c.prepend(node)

	return node.value, true
}

func (c *LRUcache[K, V]) detach(node *Node[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if c.head == node {
		c.head = node.next
	}
	if c.tail == node {
		c.tail = node.prev
	}
	node.prev = nil
	node.next = nil
}

func (c *LRUcache[K, V]) prepend(node *Node[V]) {
	if c.head == nil {
		c.head = node
		c.tail = node
	} else {
		c.head.prev = node
		node.next = c.head
		c.head = node
	}
}

func (c *LRUcache[K, V]) trimcache() {
	if c.length <= c.capacity {
		return
	}

	node := c.tail
	c.detach(c.tail)
	k := c.reverseCache[node]
	delete(c.cache, k)
	delete(c.reverseCache, node)
	c.length--
}
