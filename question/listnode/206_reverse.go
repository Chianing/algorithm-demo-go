package listnode

var ret *ListNode

// 反转链表 递归
// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
func reverseList1(head *ListNode) *ListNode {
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

// 反转链表 快慢指针
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var slowPtr, fastPtr, tmpPtr *ListNode

	for fastPtr = head; fastPtr != nil; fastPtr = tmpPtr {
		tmpPtr = fastPtr.Next
		fastPtr.Next = slowPtr
		slowPtr = fastPtr
	}

	return slowPtr
}
