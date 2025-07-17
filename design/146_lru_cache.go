type LRUCache struct {
	Capacity   int
	Cache      map[int]*Node
	Head, Tail *Node
}

type Node struct {
	Key        int
	Val        int
	Next, Prev *Node
}

func NewNode(key int, val int, next, prev *Node) *Node {
	return &Node{
		key, val, next, prev,
	}
}

func Constructor(capacity int) LRUCache {
	head, tail := NewNode(-1, 0, nil, nil), NewNode(-1, 0, nil, nil)
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		capacity,
		make(map[int]*Node, capacity),
		head,
		tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.moveToHead(node)
		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Cache[key]
	if ok {
		node.Val = value
		this.moveToHead(node)
	} else {
		node = NewNode(key, value, nil, nil)
		this.Cache[key] = node

		this.insert(node)

		if len(this.Cache) > this.Capacity {
			delete(this.Cache, this.Tail.Prev.Key)
			this.remove(this.Tail.Prev)
		}
	}
}

func (this *LRUCache) remove(node *Node) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

func (this *LRUCache) insert(node *Node) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) moveToHead(node *Node) {
	this.remove(node)
	this.insert(node)
}