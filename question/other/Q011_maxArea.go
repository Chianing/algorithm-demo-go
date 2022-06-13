package other

import "demo/util"

// 盛水最多的容器
// https://leetcode.cn/problems/container-with-most-water/
func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	ret := 0
	left, right := 0, len(height)-1

	for left < right {
		area := util.GetMinInt(height[left], height[right]) * (right - left)
		ret = util.GetMaxInt(ret, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return ret

}
