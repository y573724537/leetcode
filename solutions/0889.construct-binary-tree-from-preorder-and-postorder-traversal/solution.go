package solution

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	postIdx := findIdx(postorder, preorder[1])

	cur := &TreeNode{
		Val: preorder[0],
	}
	cur.Left = constructFromPrePost(preorder[1:postIdx+2], postorder[:postIdx+1])
	cur.Right = constructFromPrePost(preorder[postIdx+2:], postorder[postIdx+1:len(postorder)-1])

	return cur
}

func findIdx(array []int, target int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			return i
		}
	}
	return -1
}
