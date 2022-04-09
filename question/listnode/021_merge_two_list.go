package listnode

// 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	ret := &ListNode{}
	retCp := ret
	for list1 != nil || list2 != nil {
		if list1 == nil {
			retCp.Next = list2
			list2 = list2.Next
			retCp = retCp.Next
			continue
		}
		if list2 == nil {
			retCp.Next = list1
			list1 = list1.Next
			retCp = retCp.Next
			continue
		}

		if list1.Val < list2.Val {
			retCp.Next = list1
			list1 = list1.Next
		} else {
			retCp.Next = list2
			list2 = list2.Next
		}
		retCp = retCp.Next

	}

	return ret.Next

}
