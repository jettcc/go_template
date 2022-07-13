// 堆模版
package main

import "container/heap"

type Heap []int64

func (this Heap) Len() int { return len(this) }

// 实现sort.Iterface, 自定义类型需要进行修改
func (this Heap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this Heap) Less(i, j int) bool {
	return this[i] < this[j] // > 为最大堆
}

func (this *Heap) Push(v interface{}) { *this = append(*this, v.(int64)) }

func (this *Heap) Pop() interface{} {
	v := *this
	x := v[len(v)-1]
	*this = v[:len(v)-1]
	return x
}

func (this *Heap) push(v int64) { heap.Push(this, v) }

func (this *Heap) pop() int64 { return heap.Pop(this).(int64) }
