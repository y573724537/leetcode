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

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levels := [][]int{}
	queue := list.New()
	queue.PushBack(root)
	curLevel := []int{}
	curEnd := root
	var nextEnd *TreeNode

	for queue.Len() != 0 {
		cur, _ := queue.Remove(queue.Front()).(*TreeNode)
		curLevel = append(curLevel, cur.Val)

		if cur.Left != nil {
			queue.PushBack(cur.Left)
			nextEnd = cur.Left
		}

		if cur.Right != nil {
			queue.PushBack(cur.Right)
			nextEnd = cur.Right
		}

		if curEnd == cur {
			newLevels := [][]int{curLevel}
			levels = append(newLevels, levels...)
			curLevel = []int{}
			curEnd = nextEnd
		}
	}

	return levels
}

func main() {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Printf("%v\n", levelOrderBottom(root))
}
