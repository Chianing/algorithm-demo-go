package listnode

// k个一组反转链表
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/submissions/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}

	ret := &ListNode{}
	retCp := ret

	cnt := 0
	slowPtr := head

	for fastPtr := head; fastPtr != nil; {
		if cnt != k-1 {
			fastPtr = fastPtr.Next
			cnt++
			continue
		}

		tmpPtr := reverseByK1(slowPtr, k)
		retCp.Next = tmpPtr
		retCp = slowPtr
		slowPtr = slowPtr.Next
		fastPtr = slowPtr
		cnt = 0
	}

	return ret.Next
}

func reverseByK1(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cnt := 0
	var slowPtr, tmpPtr *ListNode
	fastPtr := head

	for ; cnt < k; fastPtr = tmpPtr {
		tmpPtr = fastPtr.Next
		fastPtr.Next = slowPtr
		slowPtr = fastPtr
		cnt++
	}

	head.Next = fastPtr
	return slowPtr
}
