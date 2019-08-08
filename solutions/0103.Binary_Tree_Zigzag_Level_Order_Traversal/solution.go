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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	deque := list.New()
	zigZag := [][]int{[]int{}}
	deque.PushFront(root)
	sentinel := root
	nSentinel := sentinel
	direction := true // from left to right

	for deque.Len() != 0 {
		var node *TreeNode

		node, _ = deque.Remove(deque.Back()).(*TreeNode)
		if direction {
			if node.Right != nil {
				deque.PushFront(node.Right)
				sentinel = node.Right
			}

			if node.Left != nil {
				deque.PushFront(node.Left)
				sentinel = node.Left
			}
		} else {
			if node.Left != nil {
				deque.PushFront(node.Left)
				sentinel = node.Left
			}

			if node.Right != nil {
				deque.PushFront(node.Right)
				sentinel = node.Right
			}
		}

		level := zigZag[len(zigZag)-1]
		level = append(level, node.Val)
		zigZag[len(zigZag)-1] = level

		if node == nSentinel {
			zigZag = append(zigZag, []int{})
			nSentinel = sentinel
			direction = !direction
		}
	}

	return zigZag[:len(zigZag)-1]
}
