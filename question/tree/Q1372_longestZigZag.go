package tree

import (
	"demo/util"
)

// 二叉树最长交错路径
// https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree/
func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ret := 0
	var traverse func(node *TreeNode) (leftSubPath, rightSubPath int)
	traverse = func(node *TreeNode) (leftSubPath, rightSubPath int) {
		if node == nil {
			return 0, 0
		}

		// 左子树的最长右之路径长度
		_, rightSubPath1 := traverse(node.Left)
		// 右子树的最长左之路径长度
		leftSubPath2, _ := traverse(node.Right)

		leftSubPath = 1 + rightSubPath1
		rightSubPath = 1 + leftSubPath2

		ret = util.GetMaxInt(ret, leftSubPath)
		ret = util.GetMaxInt(ret, rightSubPath)

		return leftSubPath, rightSubPath

	}
	traverse(root)

	return ret - 1

}
