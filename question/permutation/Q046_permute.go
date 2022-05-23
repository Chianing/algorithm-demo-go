package permutation

var ret [][]int
var arr []int

// 全排列
// https://leetcode.cn/problems/permutations/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	ret = make([][]int, 0)
	arr = nums

	backtrace(0)

	return ret
}

// 0~curIndex表示已选元素集合 curIndex~len(arr)-1表示候选元素集合
func backtrace(curIndex int) {
	if curIndex == len(arr)-1 {
		subRet := make([]int, len(arr))
		copy(subRet, arr)
		ret = append(ret, subRet)
		return
	}

	// 从候选元素集合选择元素
	for i := curIndex; i < len(arr); i++ {
		// 选中该元素后 将其放到已选元素集合
		swap(arr, curIndex, i)
		backtrace(curIndex + 1)
		// 撤销选择
		swap(arr, curIndex, i)
	}

}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
