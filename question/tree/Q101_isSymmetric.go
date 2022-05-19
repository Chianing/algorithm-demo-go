package tree

// 对称二叉树
// https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recur(root.Left, root.Right)

}

func recur(leftRoot, rightRoot *TreeNode) bool {
	if leftRoot == nil && rightRoot == nil {
		return true
	}

	if leftRoot == nil || rightRoot == nil || leftRoot.Val != rightRoot.Val {
		return false
	}

	return recur(leftRoot.Left, rightRoot.Right) && recur(leftRoot.Right, rightRoot.Left)
}
