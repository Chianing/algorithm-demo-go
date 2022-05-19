package tree

// 二叉树锯齿状遍历
// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *TreeNode) (ret [][]int) {
	if root == nil {
		return
	}

	que := []*TreeNode{root}
	leftFlg := true

	for len(que) > 0 {
		subRet := make([]int, 0)

		queCp := que
		que = []*TreeNode{}
		for i := 0; i < len(queCp); i++ {
			node := queCp[i]
			subRet = append(subRet, node.Val)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}

		if !leftFlg {
			for i := 0; i < len(subRet)/2; i++ {
				tmp := subRet[i]
				subRet[i] = subRet[len(subRet)-i-1]
				subRet[len(subRet)-i-1] = tmp
			}
		}

		ret = append(ret, subRet)
		leftFlg = !leftFlg
	}

	return

}
