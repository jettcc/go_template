package sort

import "fmt"

/**
 * 插入排序
 * 时间复杂度: 最好O(nlogn^2), 最坏O(n^2), 平均O(n^2)
 * 空间复杂度: O(1)
 * 稳定: false
 */
func ShellSort(nums []int) {
	gap, n := 1, len(nums)
	for gap < n/3 {
		gap *= 3
		gap++
	}

	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := nums[i]
			j := i - gap
			for j >= 0 && nums[j] > temp {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = temp
		}
		gap /= 3
	}
}
func main() {
	arr := []int{5, 7, 3, 4, 1, 2, 8, 6}
	ShellSort(arr)
	fmt.Println(arr)
}
