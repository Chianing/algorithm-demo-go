package listnode

// 两两交换链表中的节点
// https://leetcode.cn/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmpHead := &ListNode{Next: head}
	pre := tmpHead
	var cur, next *ListNode

	for pre.Next != nil && pre.Next.Next != nil {
		cur = pre.Next
		next = pre.Next.Next

		pre.Next = next
		cur.Next = next.Next
		next.Next = cur

		pre = cur

	}

	return tmpHead.Next

}
