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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levels := [][]int{}
	queue := list.New()
	queue.PushFront(root)

	for level := 0; queue.Len() != 0; level++ {
		q := queue
		queue = list.New()
		curLevel := []int{}

		for q.Len() != 0 {
			cur, _ := q.Remove(q.Back()).(*TreeNode)
			curLevel = append(curLevel, cur.Val)

			if cur.Left != nil {
				queue.PushFront(cur.Left)
			}

			if cur.Right != nil {
				queue.PushFront(cur.Right)
			}
		}

		if level%2 != 0 {
			for i, j := 0, len(curLevel)-1; i < j; {
				curLevel[i], curLevel[j] = curLevel[j], curLevel[i]
				i++
				j--
			}
		}

		levels = append(levels, curLevel)
	}

	return levels
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Printf("%v\n", zigzagLevelOrder(root))
}
