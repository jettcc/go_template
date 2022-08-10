// 队列模板
package queue

import "fmt"

type Queue struct {
	head *queueNode
	tail *queueNode
	size int
}

type queueNode struct {
	prev  *queueNode
	next  *queueNode
	value interface{}
}

func NewQueue() *Queue {
	head := &queueNode{nil, nil, nil}
	tail := &queueNode{nil, nil, nil}
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
	val := &queueNode{nil, nil, v}
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

func main() {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	queue.Push(6)
	queue.Poll()
	queue.Poll()
	queue.Poll()
	fmt.Printf("%+v\n", queue.head.next)
}
