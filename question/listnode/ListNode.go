package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 获取单节点无序链表
func getSingleListNode() *ListNode {
	return &ListNode{
		Val: 1,
	}
}

// 获取无序链表
func getListNode() *ListNode {
	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 4,
	}
	node3 := &ListNode{
		Val: 2,
	}
	node4 := &ListNode{
		Val: 5,
	}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	return node1
}

func getCustomListNode(vals []int) *ListNode {
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
