package main

import "fmt"

/**
 * 插入排序
 * 时间复杂度: 最好O(n), 最坏O(n^2), 平均O(n^2)
 * 空间复杂度: O(1)
 * 稳定: true
 */
func InsertSort(nums []int) {
	var n = len(nums)
	for i := 1; i < n; i++ {
		pre := i
		cur := nums[i]
		for pre > 0 && cur < nums[pre-1] {
			nums[pre] = nums[pre-1]
			pre--
		}
		nums[pre] = cur
	}
}

// 当数组已经是有序的，每插入一个元素，只需要考查前一个元素，因此最好情况下，插入排序的时间复杂度为O(N)。
func main() {
	arr := []int{5, 7, 3, 4, 1, 2, 8, 6}
	InsertSort(arr)
	fmt.Println(arr)
}
