package tree

// 翻转二叉树（二叉树镜像）
// https://leetcode.cn/problems/invert-binary-tree/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertTree(root.Left)
	invertTree(root.Right)

	tmpNode := root.Left
	root.Left = root.Right
	root.Right = tmpNode

	return root
}
