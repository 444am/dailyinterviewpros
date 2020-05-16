package main

type LRUCache struct {
	M    map[int]*DoubleLinkedListNode
	Cap  int
	Head *DoubleLinkedListNode
	Tail *DoubleLinkedListNode
}

type DoubleLinkedListNode struct {
	Pre  *DoubleLinkedListNode
	Next *DoubleLinkedListNode
	Key  int
	Val  int
}

func Constructor(capacity int) LRUCache {
	head := &DoubleLinkedListNode{}
	tail := &DoubleLinkedListNode{}
	head.Next = tail
	tail.Pre = head
	ret := LRUCache{M: make(map[int]*DoubleLinkedListNode), Head: head, Tail: tail, Cap: capacity}
	return ret
}

func (this *LRUCache) Get(key int) int {
	var ret int
	if v, ok := this.M[key]; ok {
		ret = v.Val
		v.Pre.Next = v.Next
		v.Next.Pre = v.Pre
		this.AppendToTail(v)
	} else {
		ret = -1
	}
	return ret
}

func (this *LRUCache) Put(key int, value int) {
	node := &DoubleLinkedListNode{Key: key, Val: value}
	if v, ok := this.M[key]; ok {
		v.Val = value
		v.Pre.Next = v.Next
		v.Next.Pre = v.Pre
		this.AppendToTail(v)
	} else {
		if len(this.M) < this.Cap {
			this.AppendToTail(node)
			this.M[key] = node
		} else {
			head := this.Head.Next
			head.Pre.Next = head.Next
			head.Next.Pre = head.Pre
			delete(this.M, head.Key)
			this.AppendToTail(node)
			this.M[key] = node
		}
	}
}

func (this *LRUCache) AppendToTail(node *DoubleLinkedListNode) {
	tail := this.Tail
	node.Pre = tail.Pre
	node.Next = tail
	tail.Pre.Next = node
	tail.Pre = node
}
