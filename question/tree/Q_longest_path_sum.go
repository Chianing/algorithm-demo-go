package tree

// 二叉树最长路径和(从根节点出发)
func longestPathSum1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLen := 0
	curLen := 0
	maxLenSum := 0
	curLenSum := 0

	var backtrace func(*TreeNode)
	backtrace = func(node *TreeNode) {
		if node == nil {
			if curLen > maxLen {
				maxLen = curLen
				maxLenSum = curLenSum
			}
			return
		}

		// 选中当前节点
		curLen++
		curLenSum += node.Val

		backtrace(node.Left)
		backtrace(node.Right)

		// 回溯
		curLen--
		curLenSum -= node.Val

	}

	backtrace(root)

	return maxLenSum

}
