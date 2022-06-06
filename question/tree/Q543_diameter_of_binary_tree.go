package tree

import "demo/util"

// 二叉树直径
// https://leetcode.cn/problems/diameter-of-binary-tree/
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ret := 0

	var getMaxDepth func(*TreeNode) int
	getMaxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := getMaxDepth(node.Left)
		rightDepth := getMaxDepth(node.Right)

		ret = util.GetMaxInt(ret, leftDepth+rightDepth)
		return util.GetMaxInt(leftDepth, rightDepth) + 1
	}

	getMaxDepth(root)

	return ret
}
