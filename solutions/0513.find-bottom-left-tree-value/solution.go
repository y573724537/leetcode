package solution

import "container/list"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)

	var node *TreeNode
	for 0 < queue.Len() {
		node, _ = queue.Remove(queue.Front()).(*TreeNode)
		if node.Right != nil {
			queue.PushBack(node.Right)
		}

		if node.Left != nil {
			queue.PushBack(node.Left)
		}
	}

	return node.Val
}
