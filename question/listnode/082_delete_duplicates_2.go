package listnode

// 删除排序链表重复元素2
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := &ListNode{}
	ret.Next = head
	retCp := ret

	for retCp.Next != nil && retCp.Next.Next != nil {
		if retCp.Next.Val != retCp.Next.Next.Val {
			retCp = retCp.Next
			continue
		}

		val := retCp.Next.Val
		for retCp.Next != nil && retCp.Next.Val == val {
			retCp.Next = retCp.Next.Next
		}
	}

	return ret.Next

}
