package tree

// 先序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)

	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

	}

	return ret
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)

	stack = append(stack, root)
	for len(stack) > 0 {
		for root.Left != nil {
			stack = append(stack, root.Left)
			root = root.Left
		}

		node := stack[len(stack)-1]
		ret = append(ret, node.Val)
		stack = stack[:len(stack)-1]

		if node.Right != nil {
			stack = append(stack, node.Right)
			root = node.Right
		}
	}

	return ret

}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root

	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node.Left != nil && node.Left != cur && node.Right != cur {
			stack = append(stack, node.Left)
		} else if node.Right != nil && node.Right != cur {
			stack = append(stack, node.Right)
		} else {
			stack = stack[:len(stack)-1]
			ret = append(ret, node.Val)
			cur = node
		}
	}

	return ret
}

// 层次遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	queue = append(queue, root)
	for len(queue) > 0 {
		subRet := make([]int, 0)
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			subRet = append(subRet, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, subRet)
	}

	return ret

}
