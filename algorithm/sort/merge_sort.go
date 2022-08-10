package sort

/**
 *归并排序
 * 时间复杂度: 最好O(nlogn), 最坏O(nlogn), 平均O(nlogn)
 * 空间复杂度: O(n)
 * 稳定: true
 */
func MergeSort(nums []int) []int {
	var n = len(nums)
	if n < 2 {
		return nums
	}
	var mid int = n >> 1
	var left []int = nums[:mid]
	var right []int = nums[mid:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	var res []int
	// 这里不能用变量表示left, right的长度，因为会缩短
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) != 0 {
		res = append(res, left...)
	}
	if len(right) != 0 {
		res = append(res, right...)
	}
	return res
}

func merge_sort(nums []int) {
	var n = len(nums)
	var temp = make([]int, n)
	sort(nums, temp, 0, n-1)
}

func sort(nums, temp []int, left, right int) {
	if left < right {
		var mid int = left + ((right - left) >> 1)
		sort(nums, temp, left, mid)
		sort(nums, temp, mid+1, right)
		t_merge(nums, temp, left, mid, right)
	}
}

func t_merge(nums, temp []int, left, mid, right int) {
	var i, j, t = left, mid + 1, 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[t] = nums[i]
			t++
			i++
		} else {
			temp[t] = nums[j]
			t++
			j++
		}
	}

	for i <= mid {
		temp[t] = nums[i]
		t++
		i++
	}

	for j <= right {
		temp[t] = nums[j]
		t++
		j++
	}
	t = 0
	for left <= right {
		nums[left] = temp[t]
		left++
		t++
	}
}
