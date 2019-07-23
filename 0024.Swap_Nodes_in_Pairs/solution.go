package solution

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	lead := &ListNode{
		Val:  0,
		Next: head,
	}

	pre, cur := lead, lead.Next

	for cur != nil && cur.Next != nil {
		pre.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, cur
		pre, cur = cur, cur.Next
	}

	return lead.Next
}
