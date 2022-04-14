package listnode

// 旋转链表
// https://leetcode-cn.com/problems/rotate-list/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}

	oldHead := head
	var oldEnd *ListNode
	length := 1

	// 找出链表长度
	for ; head.Next != nil; head = head.Next {
		length++
	}
	oldEnd = head

	// 如果k是length的倍数 不需要旋转 直接返回
	k %= length
	if k == 0 {
		return oldHead
	}

	// 首尾相接 成环
	oldEnd.Next = oldHead
	// 找出旋转后的尾节点 next指针置空
	for i := 0; i < length-k-1; i++ {
		oldHead = oldHead.Next
	}
	ret := oldHead.Next
	oldHead.Next = nil

	return ret
}
