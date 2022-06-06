package tree

import (
	"demo/util"
	"math"
)

// 二叉树中最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ret := math.MinInt

	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMaxSum := traverse(node.Left)
		rightMaxSum := traverse(node.Right)

		maxSum := node.Val + util.GetMaxInt(0, leftMaxSum) + util.GetMaxInt(0, rightMaxSum)
		ret = util.GetMaxInt(ret, maxSum)

		return node.Val + util.GetMaxInt(0, util.GetMaxInt(leftMaxSum, rightMaxSum))
	}

	traverse(root)

	return ret
}
