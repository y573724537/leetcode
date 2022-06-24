package main

import (
	"container/list"
	"fmt"
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

func largestValues(root *TreeNode) []int {
	arr := []int{}

	if root == nil {
		return arr
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		largest := math.MinInt32
		curQueue := queue
		queue = list.New()

		for curQueue.Len() != 0 {
			node, _ := curQueue.Remove(curQueue.Front()).(*TreeNode)
			if largest < node.Val {
				largest = node.Val
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		arr = append(arr, largest)
	}

	return arr
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Printf("%v", largestValues(root))
}
