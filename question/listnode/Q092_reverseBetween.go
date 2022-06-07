package listnode

// 反转链表
// https://leetcode.cn/problems/reverse-linked-list-ii/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}

	tmpHead := &ListNode{}
	tmpHead.Next = head

	leftNode := tmpHead
	for i := 0; i < left-1; i++ {
		leftNode = leftNode.Next
	}

	rightNode := leftNode
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 记录中断的节点
	subHead := leftNode.Next
	subTail := rightNode
	leftNode.Next = nil
	rightNode = rightNode.Next
	subTail.Next = nil

	// 反转截取出来的部分链表
	var reverse func(*ListNode, *ListNode)
	reverse = func(pre *ListNode, cur *ListNode) {
		if cur == nil {
			return
		}
		reverse(cur, cur.Next)
		cur.Next = pre
	}
	reverse(nil, subHead)

	// 反转后的链表拼接回去
	leftNode.Next = subTail
	subHead.Next = rightNode

	return tmpHead.Next
}

var reversedHead *ListNode

// 反转前n个元素
func reverseN(root *ListNode, n int) *ListNode {
	if n == 1 {
		reversedHead = root.Next
		return root
	}

	head := reverseN(root.Next, n-1)

	root.Next.Next = root
	root.Next = reversedHead

	return head

}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}

	head.Next = reverseBetween2(head.Next, left-1, right-1)

	return head
}
