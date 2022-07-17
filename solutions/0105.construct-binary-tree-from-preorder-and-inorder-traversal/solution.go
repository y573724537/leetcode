package solution

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	inIdx := findIdx(inorder, preorder[0])
	leftLen := inIdx

	cur := &TreeNode{
		Val: preorder[0],
	}
	cur.Left = buildTree(preorder[1:leftLen+1], inorder[:inIdx])
	cur.Right = buildTree(preorder[inIdx+1:], inorder[inIdx+1:])

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
