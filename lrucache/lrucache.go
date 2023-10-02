package lrucache

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type LRUCache[K comparable, V comparable] struct {
	Capacity     int
	Length       int
	Head         *Node[V]
	Tail         *Node[V]
	Cache        map[K]*Node[V]
	ReverseCache map[*Node[V]]K
}

func New[K comparable, V comparable](cap int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		Capacity:     cap,
		Length:       0,
		Head:         nil,
		Tail:         nil,
		Cache:        make(map[K]*Node[V]),
		ReverseCache: make(map[*Node[V]]K),
	}
}

func (c *LRUCache[K, V]) Update(key K, value V) {
	node := c.Cache[key]
	if node == nil {
		node = &Node[V]{Value: value}
		c.Length++
		c.prepend(node)
		c.trimCache()
		c.Cache[key] = node
		c.ReverseCache[node] = key
	} else {
		node.Value = value
		c.detach(node)
		c.prepend(node)
	}
}

func (c *LRUCache[K, V]) Get(key K) (val V, ok bool) {
	node := c.Cache[key]
	if node == nil {
		return val, false
	}

	c.detach(node)
	c.prepend(node)

	return node.Value, true
}

func (c *LRUCache[K, V]) detach(node *Node[V]) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if c.Head == node {
		c.Head = node.Next
	}
	if c.Tail == node {
		c.Tail = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}

func (c *LRUCache[K, V]) prepend(node *Node[V]) {
	if c.Head == nil {
		c.Head = node
		c.Tail = node
	} else {
		c.Head.Prev = node
		node.Next = c.Head
		c.Head = node
	}
}

func (c *LRUCache[K, V]) trimCache() {
	if c.Length <= c.Capacity {
		return
	}

	node := c.Tail
	c.detach(c.Tail)
	k := c.ReverseCache[node]
	delete(c.Cache, k)
	delete(c.ReverseCache, node)
	c.Length--
}
