package algorithm

// 二分查找
// https://leetcode-cn.com/problems/binary-search/
func binarySearch(nums []int, target int) int {
	length := len(nums)
	if length == 0 || target < nums[0] || target > nums[length-1] {
		return -1
	}

	start := 0
	end := length - 1

	for start <= end {
		mid := (start + end) / 2
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1

}
