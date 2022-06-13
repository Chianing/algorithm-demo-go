package tree

// 合并二叉树
// https://leetcode.cn/problems/merge-two-binary-trees/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	var merge func(*TreeNode, *TreeNode) *TreeNode
	merge = func(node1 *TreeNode, node2 *TreeNode) *TreeNode {
		if node1 == nil {
			return node2
		}
		if node2 == nil {
			return node1
		}

		node1.Val = node1.Val + node2.Val
		node1.Left = merge(node1.Left, node2.Left)
		node1.Right = merge(node1.Right, node2.Right)

		return node1
	}
	merge(root1, root2)

	return root1

}
