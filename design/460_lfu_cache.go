type LFUCache struct {
	Capacity int
	MinFreq  int
	Nodes    map[int]*Node
	FreqMap  map[int]*LRUCache // key is frequency
}

type LRUCache struct {
	Head, Tail *Node
}

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
	Freq  int
}

func NewLRUCache() *LRUCache {
	head, tail := NewNode(0, 0), NewNode(0, 0)
	head.Next = tail
	tail.Prev = head

	return &LRUCache{
		Head: head,
		Tail: tail,
	}
}

// check if lru is empty
func (c *LRUCache) isEmpty() bool {
	return c.Head.Next == c.Tail
}

func (c *LRUCache) addToFront(node *Node) {
	node.Next = c.Head.Next
	node.Prev = c.Head
	c.Head.Next.Prev = node
	c.Head.Next = node
}

// works only when lru not empty, check is empty before
func (c *LRUCache) popLRU() *Node {
	lruNode := c.Tail.Prev

	lruNode.remove()

	return lruNode
}

// default freq is 1
func NewNode(key int, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
		Freq:  1,
	}
}

func (node *Node) remove() {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Capacity: capacity,
		MinFreq:  1,
		Nodes:    make(map[int]*Node),
		FreqMap:  make(map[int]*LRUCache),
	}
}

func (c *LFUCache) Get(key int) int {
	if node, ok := c.Nodes[key]; ok {
		c.increaseFreq(node)
		return node.Value
	}

	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if node, ok := c.Nodes[key]; ok { // update
		node.Value = value
		c.increaseFreq(node)
	} else { // push
		if len(c.Nodes) == c.Capacity { // reach limit
			victimLRU := c.FreqMap[c.MinFreq]
			victimNode := victimLRU.popLRU()
			delete(c.Nodes, victimNode.Key)

			if victimLRU.isEmpty() {
				delete(c.FreqMap, c.MinFreq)

				if victimLRU.isEmpty() {
					delete(c.FreqMap, c.MinFreq)
				}
			}
		}

		c.MinFreq = 1
		newNode := NewNode(key, value)

		c.Nodes[key] = newNode
		c.ensureLRU(c.MinFreq).addToFront(newNode)
	}
}

func (c *LFUCache) increaseFreq(node *Node) {
	oldFreq, newFreq := node.Freq, node.Freq+1
	node.Freq = newFreq

	// remove from old lru
	node.remove()
	if c.FreqMap[oldFreq].isEmpty() {
		delete(c.FreqMap, oldFreq)
		if oldFreq == c.MinFreq {
			c.MinFreq = newFreq
		}
	}

	// insert to new lru
	c.ensureLRU(newFreq).addToFront(node)
}

func (c *LFUCache) ensureLRU(freq int) *LRUCache {
	cache := c.FreqMap[freq]
	if cache == nil {
		cache = NewLRUCache()
		c.FreqMap[freq] = cache
	}
	return cache
}
