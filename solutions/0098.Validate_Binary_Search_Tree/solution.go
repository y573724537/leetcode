package solution

import (
	"container/list"
	"math"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	cur := root
	prev := math.MinInt64
	stack := list.New()

	for cur != nil || stack.Len() != 0 {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}

		node, _ := stack.Remove(stack.Back()).(*TreeNode)
		if node.Val <= prev {
			return false
		}

		prev = node.Val
		cur = node.Right
	}

	return true
}
