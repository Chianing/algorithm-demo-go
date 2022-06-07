package tree

// 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ret := make([][]int, 0)
	que := []*TreeNode{root}

	for len(que) > 0 {
		size := len(que)
		subRet := make([]int, size)

		for i := 0; i < size; i++ {
			node := que[i]
			subRet[i] = node.Val
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}

		ret = append(ret, subRet)
		que = que[size:]

	}

	var swap func(int, int)
	swap = func(i1 int, i2 int) {
		tmp := ret[i1]
		ret[i1] = ret[i2]
		ret[i2] = tmp
	}

	var reverseList func()
	reverseList = func() {
		for i := 0; i < len(ret)/2; i++ {
			swap(i, len(ret)-i-1)
		}
	}

	reverseList()

	return ret

}
