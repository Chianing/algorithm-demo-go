package tree

import "demo/util"

// 二叉树的坡度
// https://leetcode.cn/problems/binary-tree-tilt/
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ret := 0
	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := traverse(node.Left)
		right := traverse(node.Right)

		ret += util.GetAbsInt(left - right)

		return node.Val + left + right
	}

	traverse(root)

	return ret

}
