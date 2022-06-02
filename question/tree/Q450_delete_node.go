package tree

// 删除二叉搜索树中节点
// https://leetcode.cn/problems/delete-node-in-a-bst/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	// 要删除的节点在当前节点的右子树中
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	// 要删除的节点在当前节点的左子树中
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	// 当前节点就是要删除的节点
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	// 当前节点的左节点将要被提上来 所以将当前节点右子树放到左节点最右子节点右面
	tmpNode := root.Left
	for tmpNode.Right != nil {
		tmpNode = tmpNode.Right
	}
	tmpNode.Right = root.Right

	return root.Left

}
