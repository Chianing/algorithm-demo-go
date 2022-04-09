package listnode

// 两数相加
// https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	var tmp, sum int
	var ret = &ListNode{}
	retCp := ret

	for l1 != nil || l2 != nil {
		if l1 == nil {
			sum = l2.Val + tmp
			l2 = l2.Next
		} else if l2 == nil {
			sum = l1.Val + tmp
			l1 = l1.Next
		} else {
			sum = l1.Val + l2.Val + tmp
			l1 = l1.Next
			l2 = l2.Next
		}

		retCp.Next = &ListNode{
			Val: sum % 10,
		}
		tmp = sum / 10
		retCp = retCp.Next
	}

	if tmp != 0 {
		retCp.Next = &ListNode{
			Val: tmp % 10,
		}
	}

	return ret.Next
}
