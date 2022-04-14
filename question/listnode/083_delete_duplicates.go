package listnode

// 删除有序链表中重复元素
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := head

	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return ret
}
