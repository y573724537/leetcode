package solution

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	prev := head
	carry, curVal := 0, 0

	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		curVal, carry = val%10, val/10
		cur := &ListNode{Val: curVal}
		prev.Next, prev, l1, l2 = cur, cur, l1.Next, l2.Next
	}

	for l1 != nil {
		val := l1.Val + carry
		curVal, carry = val%10, val/10
		cur := &ListNode{Val: curVal}
		prev.Next, prev, l1 = cur, cur, l1.Next
	}

	for l2 != nil {
		val := l2.Val + carry
		curVal, carry = val%10, val/10
		cur := &ListNode{Val: curVal}
		prev.Next, prev, l2 = cur, cur, l2.Next
	}

	if carry != 0 {
		prev.Next = &ListNode{Val: carry}
	}

	return head.Next
}
