package listnode

var ret *ListNode

// 反转链表
// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	subReverseList(nil, head)

	return ret
}

func subReverseList(pre *ListNode, node *ListNode) {
	if node == nil {
		ret = pre
		return
	}
	subReverseList(node, node.Next)
	node.Next = pre
}
