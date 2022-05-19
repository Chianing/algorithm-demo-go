package tree

// 二叉树是否互为镜像
func isMirror(root1, root2 *TreeNode) bool {
	if root1 == root2 {
		return true
	}

	if (root1 == nil && root2 != nil) ||
		(root1 != nil && root2 == nil) ||
		(root1.Val != root2.Val) {
		return false
	}

	return isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)

}
