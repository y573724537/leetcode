package main

import (
	"fmt"
)

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

func sortList(head *ListNode) *ListNode {
	return mergeSortList(head, nil)
}

// tail is sentinel
func mergeSortList(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		fast = fast.Next
		slow = slow.Next

		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow

	part1 := mergeSortList(head, mid)
	part2 := mergeSortList(mid, tail)

	return merge(part1, part2)
}

func merge(part1, part2 *ListNode) *ListNode {
	head := new(ListNode)
	cur, p1, p2 := head, part1, part2

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}

		cur = cur.Next
	}

	if p1 != nil {
		cur.Next = p1
	}

	if p2 != nil {
		cur.Next = p2
	}

	return head.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
	}

	l0 := &ListNode{
		Val:  4,
		Next: l1,
	}

	fmt.Printf("%v\n", sortList(l0))
}
