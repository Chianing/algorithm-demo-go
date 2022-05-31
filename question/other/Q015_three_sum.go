package other

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{}
	}

	sort.Ints(nums)
	length := len(nums)
	ret := make([][]int, 0)

	for cur, left, right := 0, 0, 0; cur < length-2; cur++ {
		// 排序后如果当前数值大于0 说明后面的元素都大于0 则不可能存在三个正数之和为0
		if nums[cur] > 0 {
			return ret
		}
		// 避免重复
		if cur > 0 && nums[cur] == nums[cur-1] {
			continue
		}

		left = cur + 1
		right = length - 1

		for left < right {
			sum := nums[cur] + nums[left] + nums[right]

			if sum > 0 {
				right--
				continue
			}
			if sum < 0 {
				left++
				continue
			}

			ret = append(ret, []int{nums[cur], nums[left], nums[right]})
			// 去重 如果相等 证明以该元素为结果的集已经存在了
			for left < right && nums[left] == nums[left+1] {
				left++
			}
			for left < right && nums[right] == nums[right-1] {
				right--
			}
			right--
			left++
		}

	}

	return ret

}
