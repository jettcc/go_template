package sort

/**
 * 选择排序
 * 时间复杂度: 最好O(n^2), 最坏O(n^2), 平均O(n^2)
 * 空间复杂度: O(1)
 * 稳定: false
 */

func SelectSort(nums []int) {
	var n = len(nums)
	for i := 0; i < n; i++ {
		var min = i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
}
