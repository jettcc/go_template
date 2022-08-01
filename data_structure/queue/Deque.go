// 双端队列模板
package main

type Deque struct {
	head *node
	tail *node
	size int
}

type node struct {
	prev  *node
	next  *node
	value interface{}
}

func NewDeque() *Deque {
	head := &node{nil, nil, nil}
	tail := &node{nil, nil, nil}
	head.next = tail
	tail.prev = head
	return &Deque{head, tail, 0}
}

func (this *Deque) Size() int {
	return this.size
}

func (this *Deque) IsEmpty() bool {
	return this.size == 0
}

func (this *Deque) Peek() interface{} {
	if this.head == nil {
		return nil
	}
	return this.head.value
}

func (this *Deque) PushFirst(v interface{}) {
	val := &node{nil, nil, v}
	last := this.head.next

	this.head.next = val
	val.next = last
	val.prev = this.head
	last.prev = val

	this.size++
}

func (this *Deque) PushLast(v interface{}) {
	val := &node{nil, nil, v}
	front := this.tail.prev

	front.next = val
	val.prev = front

	this.tail.prev = val
	val.next = this.tail
	this.size++
}

func (this *Deque) PollFirst() interface{} {
	if this.size == 0 {
		return nil
	}

	val := this.head.next
	this.head.next = val.next
	val.next.prev = this.head
	this.size--
	return val.value
}

func (this *Deque) pollLast() interface{} {
	if this.size == 0 {
		return nil
	}
	val := this.tail.prev
	this.tail.prev = val.prev
	val.prev.next = this.tail
	this.size--
	return val.value
}
