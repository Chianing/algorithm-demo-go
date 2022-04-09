package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 获取单节点无序链表
func getSingleListNode() *ListNode {
	return getCustomListNode(1)
}

// 获取无序链表
func getListNode() *ListNode {
	return getCustomListNode(1, 3, 4, 2, 5)
}

func getCustomListNode(vals ...int) *ListNode {
	ret := &ListNode{}
	retCp := ret
	for _, val := range vals {
		retCp.Next = &ListNode{
			Val: val,
		}
		retCp = retCp.Next
	}

	return ret.Next
}
