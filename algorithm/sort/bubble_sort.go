package main

/**
 * 冒泡排序
 * 时间复杂度: 最好O(n), 最坏O(n^2), 平均O(n^2)
 * 空间复杂度: O(1)
 * 稳定: true
 */
func BubbleSort(nums []int) {
	var n = len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
