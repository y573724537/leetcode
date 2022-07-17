package solution

import (
	"container/list"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levels := [][]int{}
	queue := list.New()
	queue.PushFront(root)
	curLevel := []int{}
	curEnd := root
	var nextEnd *TreeNode

	for queue.Len() != 0 {
		node, _ := queue.Remove(queue.Back()).(*TreeNode)
		curLevel = append(curLevel, node.Val)

		if node.Left != nil {
			queue.PushFront(node.Left)
			nextEnd = node.Left
		}

		if node.Right != nil {
			queue.PushFront(node.Right)
			nextEnd = node.Right
		}

		if node == curEnd {
			levels = append(levels, curLevel)
			curLevel = []int{}
			curEnd = nextEnd
		}
	}

	return levels
}

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levels := [][]int{}
	queue := list.New()
	queue.PushFront(root)

	for queue.Len() != 0 {
		q := queue
		queue = list.New()
		level := []int{}

		for q.Len() != 0 {
			cur, _ := q.Remove(q.Back()).(*TreeNode)
			level = append(level, cur.Val)

			if cur.Left != nil {
				queue.PushFront(cur.Left)
			}

			if cur.Right != nil {
				queue.PushFront(cur.Right)
			}
		}

		levels = append(levels, level)
	}

	return levels
}
