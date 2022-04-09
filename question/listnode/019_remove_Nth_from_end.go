package listnode

// 删除链表倒数第n个节点
// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/submissions/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	slowPtrPre := head
	slowPtr := head
	fastPtr := head
	diffCount := 0

	for ; fastPtr.Next != nil; fastPtr = fastPtr.Next {
		if diffCount < n-1 {
			diffCount++
			continue
		}
		slowPtrPre = slowPtr
		slowPtr = slowPtr.Next
	}

	// n大于链表长度
	if diffCount != n-1 {
		return head
	}
	// 删除的是第一个节点
	if slowPtr == slowPtrPre {
		return head.Next
	}
	// 删除倒数第n个节点
	slowPtrPre.Next = slowPtr.Next

	return head
}
