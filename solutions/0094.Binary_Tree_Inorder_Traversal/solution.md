## 二叉树遍历

### 递归遍历

1. 先序遍历
```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    return preorderTraverse(root, nil)
}

func preorderTraverse(root *TreeNode, preorder []int) []int {
    if preorder == nil {
        preorder = []int{}
    }

    if root == nil {
        return preorder
    }

    preorder = append(preorder, root.Val)
    preorder = preorderTraverse(root.Left, preorder))
    preorder = preorderTraverse(root.Right, preorder)

    return preorder
}
```

2. 中序遍历
```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    return inorderTraverse(root, nil)
}

func inorderTraverse(root *TreeNode, inorder []int) []int {
    if inorder == nil {
        inorder = []int{}
    }

    if root == nil {
        return inorder
    }

    inorder = inorderTraverse(root.Left, inorder)
    inorder = append(inorder, root.Val)
    inorder = inorderTraverse(root.Right, inorder)

    return inorder
}
```

3. 后序遍历
```golang
func postorderTraversal(root *TreeNode) []int {
    return postorderTraverse(root, nil)
}

func postorderTraverse(root *TreeNode, postorder []int) []int {
    if postorder == nil {
        postorder = []int{}
    }

    if root == nil {
        return postorder
    }

    postorder = postorderTraverse(root.Left, postorder)
    postorder = postorderTraverse(root.Right, postorder)
    postorder = append(postorder, root.Val)

    return postorder
}
```

### 非递归遍历
1. 先序遍历
```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    preorder := []int{}

    if root == nil {
        return preorder
    }

    stack := list.New()
    stack.PushBack(root)

    for stack.Len() > 0 {
        node, _ := stack.Remove(stack.Back()).(*TreeNode)
        preorder = append(preorder, root.Val)

        if node.Right != nil {
            stack.PushBack(node.Right)
        }

        if node.Left != nil {
            stack.PushBack(node.Left)
        }
    }

    return preorder
}
```

2. 中序遍历
```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    inorder := []int{}

    if root == nil {
        return inorder
    }

    stack := list.New()
    cur := root
    for cur != nil || stack.Len() > 0 {
        for cur != nil {
            stack.PushBack(cur)
            cur = cur.Left
        }

        node, _ := stack.Remove(stack.Back()).(*TreeNode)
        inorder = append(inorder, node.Val)

        cur = node.Right
    }

    return inorder
}
```

3. 后序遍历
```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    postorder := []int{}

    if root == nil {
        return postorder
    }

    s1 := list.New()
    s2 := list.New()

    s1.PushBack(root)
    for s1.Len() > 0 {
        node, _ := s1.Remove(s1.Back()).(*TreeNode)
        s2.PushBack(node)

        if node.Left != nil {
            s1.PushBack(node.Left)
        }

        if node.Right != nil {
            s1.PushBack(node.Right)
        }
    }

    for s2.Len() > 0 {
        node, _ := s2.Remove(s1.Back()).(*TreeNode)
        postorder = append(postorder, node.Val)
    }

    return postorder
}
```

















