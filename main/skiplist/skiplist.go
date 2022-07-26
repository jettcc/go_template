package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SkipList struct {
	head       *Node
	size       int
	level      int  // 当前层级: 指的是当前跳表中有多少层
	repetition bool // 是否允许重复
}

type Node struct {
	Score int64
	value interface{} // 实践中可把int64改成其他数据类型
	next  []*Node
}

const (
	maxLevel   = 16 // 最大层级: 指的是跳表最多能有多少层
	p          = 0.25
	ERROR_CODE = -43962200
)

func newNode(Score int64, value interface{}, maxLevel int) *Node {
	return &Node{
		Score: Score,
		value: value,
		next:  make([]*Node, maxLevel),
	}
}

func New(repetition bool) *SkipList {
	return &SkipList{
		head:       &Node{next: make([]*Node, maxLevel)},
		repetition: repetition,
	}
}

/**
 * 获取随机level
 * 这里参考了skiplist论文中的获取随机level的方法
 */
func randomLevel() int {
	var level = 1
	for level < maxLevel && rands() < p {
		level++
	}
	return level
}

func rands() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}

/**
 * 搜索算法
 */
func (this *SkipList) serch(searchScore int64) (*Node, bool) {
	// 从最高层开始从左向右找（最上面一层可以最快定位到我们想找的元素大概位置）
	c := this.head
	for i := this.level - 1; i >= 0; i-- {
		// 退出循环的条件就是next元素大于指定的元素, 就往下一层开始找
		for c.next[i] != nil && c.next[i].Score < searchScore {
			// 当前节点在这一层直接向前
			c = c.next[i]
		}
		// 判断下一个元素是否满足条件, 任何一层, 找到就直接返回对应的值
		if c.next[i] != nil && c.next[i].Score == searchScore {
			return c.next[i], true
		}
	}
	// 找到最下面一层还没有找到则说明元素不存在。
	return nil, false
}

/**
 * 插入算法
 */
func (this *SkipList) Insert(score int64, value interface{}) {
	// update用来记录每一层最后一个小于score的节点
	update := make([]*Node, maxLevel)
	c := this.head
	for i := this.level - 1; i >= 0; i-- {
		for c.next[i] != nil && c.next[i].Score < score {
			// 当前节点在这一层直接向前
			c = c.next[i]
		}
		update[i] = c
	}
	c = c.next[0]
	// 如果这个元素本身就存在跳表里, 则直接更新
	// 如果允许跳表中可以包含重复元素, 这里注释掉即可
	if c != nil && !this.repetition && c.Score == score {
		c.value = value
		return
	}
	/** 如果随机到的level比当前的level要高, 就更新层级
	 * 一种实现方式是可以把rdmlevel设置current level + 1, 方便查找
	 * 但是破坏了算法的随机性
	 * // The pseudo-code
	 * if level > this.level {
	 * 		level = this.level + 1
	 *		update[this.level] = this.header
	 *		this.level = level
	 *	}
	 */
	rdmlevel := randomLevel()
	if rdmlevel > this.level {
		for i := this.level; i < rdmlevel; i++ {
			update[i] = this.head
		}
		this.level = rdmlevel
	}
	e := newNode(score, value, rdmlevel)
	// 处理每一层的情况
	// o->last => o->update->last
	for i := 0; i < rdmlevel; i++ {
		e.next[i] = update[i].next[i]
		update[i].next[i] = e
	}
	this.size++
}

/**
 * 删除算法
 */
func (this *SkipList) delete(searchScore int64) bool {
	c := this.head
	update := make([]*Node, maxLevel)
	for i := this.level - 1; i >= 0; i-- {
		for c.next[i] != nil && c.next[i].Score < searchScore {
			c = c.next[i]
		}
		// 设置每一层对应的元素信息
		update[i] = c
	}
	// 最下面一层的第一个指向的元素, 即即将被删除的元素
	c = c.next[0]

	if c != nil && c.Score == searchScore {
		for i := 0; i < this.level; i++ {
			// 处理每一层的情况
			// o->update->last => o->last
			if update[i].next[i] != c {
				break
			}
			update[i].next[i] = c.next[i]
		}
		this.size--
		return true
	}
	return false
}

/*
	main
*/
func main() {
	sk := New(false)
	sk.Insert(0, 3)
	sk.Insert(0, 4)
	for _, row := range sk.head.next {
		if row != nil {
			fmt.Printf("%+v\n", row)
		}
	}

}
