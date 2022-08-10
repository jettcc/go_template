package queue

type Item struct {
	value    int64 // 随意类型
	priority int   // 优先级
	index    int   // 索引
}

type PriorityQueue []*Item

func (this PriorityQueue) Len() int           { return len(this) }
func (this PriorityQueue) Less(i, j int) bool { return this[i].priority > this[j].priority }
func (this PriorityQueue) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *PriorityQueue) Push(x interface{}) {
	n := len(*this)
	item := x.(*Item)
	item.index = n
	*this = append(*this, item)
}

func (this *PriorityQueue) Pop() interface{} {
	old := *this
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*this = old[0 : n-1]
	return item
}
