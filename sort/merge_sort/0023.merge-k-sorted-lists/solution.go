package main

import "fmt"

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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) < 2 {
		return lists[0]
	}

	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, low, high int) *ListNode {
	if low == high {
		return lists[low]
	}

	mid := low + (high-low)>>1

	return mergeTowLists(merge(lists, low, mid), merge(lists, mid+1, high))
}

func mergeTowLists(part1, part2 *ListNode) *ListNode {
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
	} else if p2 != nil {
		cur.Next = p2
	}

	return head.Next
}

func main() {
	lists := []*ListNode{
		nil,
		{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 11}}},
		nil,
		{Val: 6, Next: &ListNode{Val: 10}},
	}
	fmt.Printf("%v\n", mergeKLists(lists))
}
