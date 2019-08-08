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

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	sequence := []int{}
	cur := root
	stack := list.New()

	for cur != nil || stack.Len() != 0 {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}

		node, _ := stack.Remove(stack.Back()).(*TreeNode)
		sequence = append(sequence, node.Val)

		cur = node.Right
	}

	return sequence
}
