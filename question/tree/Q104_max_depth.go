package tree

// 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return getMaxInt(leftDepth, rightDepth) + 1

}

func getMaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
