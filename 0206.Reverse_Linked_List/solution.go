package solution

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre *ListNode
	for cur := head; cur != nil; {
		cur.Next, cur, pre = pre, cur.Next, cur
	}

	return pre.Next
}
