package main

import (
	"container/list"
	"fmt"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	sequence := []int{}
	stack1 := list.New()
	stack2 := list.New()
	stack1.PushBack(root)

	for stack1.Len() != 0 {
		node, _ := stack1.Remove(stack1.Back()).(*TreeNode)
		stack2.PushBack(node)

		if node.Left != nil {
			stack1.PushBack(node.Left)
		}

		if node.Right != nil {
			stack1.PushBack(node.Right)
		}
	}

	for stack2.Len() != 0 {
		node, _ := stack2.Remove(stack2.Back()).(*TreeNode)
		sequence = append(sequence, node.Val)
	}

	return sequence
}

func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	sequence := []int{}
	stack := list.New()
	stack.PushBack(root)
	var cur *TreeNode
	var last *TreeNode = root

	for stack.Len() != 0 {
		cur, _ = stack.Back().Value.(*TreeNode)
		if cur.Left != nil && last != cur.Left && last != cur.Right {
			stack.PushBack(cur.Left)
		} else if cur.Right != nil && last != cur.Right {
			stack.PushBack(cur.Right)
		} else {
			node, _ := stack.Remove(stack.Back()).(*TreeNode)
			sequence = append(sequence, node.Val)
			last = cur
		}
	}

	return sequence
}

func main() {
	t3 := &TreeNode{3, nil, nil}
	t2 := &TreeNode{2, t3, nil}
	t1 := &TreeNode{1, nil, t2}

	fmt.Printf("%v\n", postorderTraversal(t1))
}
