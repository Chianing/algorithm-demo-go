package tree

import "demo/util"

// 平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	balancedFlg := true
	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftDepth := traverse(node.Left)
		rightDepth := traverse(node.Right)

		if util.GetAbsInt(leftDepth-rightDepth) > 1 {
			balancedFlg = false
		}

		return util.GetMaxInt(leftDepth, rightDepth) + 1

	}
	traverse(root)

	return balancedFlg
}
