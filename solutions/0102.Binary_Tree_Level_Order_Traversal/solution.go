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
		return nil
	}

	levels := [][]int{[]int{}}
	queue := list.New()
	queue.PushFront(root)
	sentinel := root
	nSentinel := sentinel

	for queue.Len() != 0 {
		node, _ := queue.Remove(queue.Back()).(*TreeNode)

		level := levels[len(levels)-1]
		level = append(level, node.Val)
		levels[len(levels)-1] = level

		if node.Left != nil {
			queue.PushFront(node.Left)
			sentinel = node.Left
		}

		if node.Right != nil {
			queue.PushFront(node.Right)
			sentinel = node.Right
		}

		if node == nSentinel {
			levels = append(levels, []int{})
			nSentinel = sentinel
		}
	}

	return levels[:len(levels)-1]
}
