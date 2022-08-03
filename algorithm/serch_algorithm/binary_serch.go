package main

/**
 * 基本二分, 用于查找指定元素的下标, 如果不存在返回 -1
 * 查找条件可以在不与元素的两侧进行比较的情况下确定（或使用它周围的特定元素）。
 * 不需要后处理，因为每一步中，你都在检查是否找到了元素。如果到达末尾，则知道未找到该元素。
 */
func binarySearch(nums []int, tar int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		var mid = l + ((r - l) >> 1)
		if nums[mid] == tar {
			return mid
		} else if nums[mid] > tar {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

/**
 * 寻找左侧边界的二分查找
 */
func left_binarySerch(nums []int, tar int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, n
	for l < r {
		var mid = l + ((r - l) >> 1)
		if nums[mid] < tar {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

/**
 * 寻找右侧边界的二分查找
 */
func right_binarySerch(nums []int, tar int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, n
	for l < r {
		var mid = l + ((r - l) >> 1)
		if nums[mid] > tar {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r - 1
}
