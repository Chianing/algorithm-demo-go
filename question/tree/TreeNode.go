package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
  3
 / \
9  20
  /  \
 15   7
*/
func getTreeNode() *TreeNode {
	ret := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}

	ret.Left = node1
	ret.Right = node2
	node2.Left = node3
	node2.Right = node4

	return ret
}

/*
    3
   / \
  20  9
 /  \
7   15
*/
func getMirrorTreeNode() *TreeNode {
	ret := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 20}
	node2 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 15}

	ret.Left = node1
	ret.Right = node2
	node1.Left = node3
	node1.Right = node4

	return ret
}
