package solution

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	inIdx := findIdx(inorder, postorder[len(postorder)-1])

	cur := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	cur.Left = buildTree(inorder[:inIdx], postorder[:inIdx])
	cur.Right = buildTree(inorder[inIdx+1:], postorder[inIdx:len(postorder)-1])

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
