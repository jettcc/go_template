// 队列模板
package main

type Queue struct {
	head *node
	tail *node
	size int
}

type node struct {
	prev  *node
	next  *node
	value interface{}
}

func NewQueue() *Queue {
	head := &node{nil, nil, nil}
	tail := &node{nil, nil, nil}
	head.next = tail
	tail.prev = head
	return &Queue{head, tail, 0}
}

func (this *Queue) Size() int {
	return this.size
}

func (this *Queue) IsEmpty() bool {
	return this.size == 0
}

func (this *Queue) Peek() interface{} {
	if this.head == nil {
		return nil
	}
	return this.head.value
}

func (this *Queue) Push(v interface{}) {
	val := &node{nil, nil, v}
	front := this.tail.prev

	front.next = val
	val.prev = front

	this.tail.prev = val
	val.next = this.tail
	this.size++
}

func (this *Queue) Poll() interface{} {
	if this.size == 0 {
		return nil
	}

	val := this.head.next
	this.head.next = val.next
	val.next.prev = this.head
	this.size--
	return val.value
}
