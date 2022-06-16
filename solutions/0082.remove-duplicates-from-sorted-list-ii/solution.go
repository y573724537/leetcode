package solution

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	fakeHead := &ListNode{
		Next: head,
	}

	cur, compare := fakeHead, head
	for compare != nil && compare.Next != nil {
		next, compareVal := compare, compare.Val

		for next != nil && next.Next != nil && compareVal == next.Next.Val {
			next = next.Next
		}

		if next == compare {
			cur = compare
		} else {
			cur.Next = next.Next
		}
		compare = cur.Next
	}

	return fakeHead.Next
}
