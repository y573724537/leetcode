package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	lastSorted, cur := head, head.Next

	for cur != nil {
		if lastSorted.Val <= cur.Val {
			lastSorted, cur = cur, cur.Next
			continue
		}

		pre := dummyHead
		for pre.Next.Val <= cur.Val {
			pre = pre.Next
		}

		lastSorted.Next = cur.Next
		pre.Next, cur.Next = cur, pre.Next
		cur = lastSorted.Next
	}

	return dummyHead.Next
}
