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

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	sequence := []int{}
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() != 0 {
		node, _ := stack.Remove(stack.Back()).(*TreeNode)
		sequence = append(sequence, node.Val)

		if node.Right != nil {
			stack.PushBack(node.Right)
		}

		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}

	return sequence
}
