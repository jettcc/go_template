// 栈模板
type Stack struct {
	head *node
	tail *node
	size int
}

type node struct {
	prev  *node
	next  *node
	value interface{}
}

func NewStack() *Stack {
	head := &node{nil, nil, nil}
	tail := &node{nil, nil, nil}
	head.next = tail
	tail.prev = head
	return &Stack{head, tail, 0}
}

func (this *Stack) Size() int {
	return this.size
}

func (this *Stack) IsEmpty() bool {
	return this.size == 0
}

func (this *Stack) Peek() interface{} {
	if this.head == nil {
		return nil
	}
	return this.head.value
}

func (this *Stack) Push(v interface{}) {
	if v == nil {
		return
	}
	val := &node{nil, nil, v}
	node := this.head.next

	node.prev = val
	this.head.next = val

	val.next = node
	val.prev = this.head
	this.size++
}

func (this *Stack) Pop() interface{} {
	if this.size == 0 {
		return nil
	}

	node := this.head.next
	this.head.next = node.next
	node.next.prev = this.head
	this.size--
	return node.value
}
