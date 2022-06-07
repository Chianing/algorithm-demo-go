package tree

// 二叉树右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)
	que := []*TreeNode{root}

	for len(que) > 0 {
		size := len(que)

		for i := 0; i < size; i++ {
			node := que[i]
			if i == size-1 {
				ret = append(ret, que[i].Val)
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}

		}

		que = que[size:]
	}

	return ret

}
