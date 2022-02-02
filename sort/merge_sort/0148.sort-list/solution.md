# 思路

## 自顶向下归并排序
题目要求时间复杂度为O(nlogn)，可以想到使用归并排序，不过自顶向下的归并空间复杂度为O(logn)，空间复杂度主要取决于递归调用的栈空间。数组归并也会有该空间的开销，不过由于拷贝数组时还会开辟O(n)的空间，因此将该开销忽略了。

该题有不少比较tricky的地方，问题单独拿出来，都比较好解，一旦合并到该题内，就比较麻烦

### 获取链表中点
leetcode原题: [876. Middle of the Linked List](https://leetcode-cn.com/problems/middle-of-the-linked-list/)

典型的双指针问题

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }

    return slow
}
```

不过运用在本题中，则需要在链表中起点以及终点之间获取中间节点，注意head是起点，tail不是终点，而是链表的sentinel
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head, tail *ListNode) *ListNode {
    if head == nil {
        return head
    }

    if head.Next = tail {
        return head
    }

    slow, fast := head, head

    for fast != tail {
        fast = fast.Next
        slow = slow.Next

        if fast != tail {
            fast = fast.Next
        }
    }

    return slow
}
```

### 合并两个有序链表
leetcode原题: [21. Merge Two Sorted Lists](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

也比较简单，但问题是，需要合并两个独立的链表，确保两个链表都有结尾sentinel为空
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    p1, p2 := list1, list2
    head := new(ListNode)
    cur := head

    for p1 != nil && p2 != nil {
        if p1.Val <= p2.Val {
            cur.Next = p1
            p1 = p1.Next
        } else {
            cur.Next = p2
            p2 = p2.Next
        }

        cur = cur.Next
    }

    if p1 != nil {
        cur.Next = p1
    }

    if p2 != nil {
        cur.Next = p2
    }

    return head.Next
}
```

### 基于上述基础算法，设计链表归并排序

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    // 对于起点head，到sentinel nil进行递归
    return mergeSort(head, nil)
}

func mergeSort(head, sentinel *ListNode) *ListNode {
    // 没有节点，不需要排序
    if head == nil {
        return head
    }

    // 仅有一个节点，不需要排序
    if head.Next == sentinel {
        // 为了保证merge时，两个链表独立，一定要注意！！！
        head.Next = nil
        return head
    }

    // 获取中点
    slow, fast := head, head
    for fast != sentinel {
        fast = fast.Next
        slow = slow.Next

        if fast != sentinel {
            fast = fast.Next
        }
    }

    mid := slow

    return merge(mergeSort(head, mid), mergeSort(mid, sentinel))
}

func merge(part1, part2 *ListNode) *ListNode {
    head := new(ListNode)
    cur, p1, p2 := head, part1, part2

    for p1 != nil && p2 != nil {
        if p1.Val <= p2.Val {
            cur.Next = p1
            p1 = p1.Next
        } else {
            cur.Next = p2
            p2 = p2.Next
        }
        cur = cur.Next
    }

    // 没有必要像数组一样，逐个拷贝了，指针挂上即可
    if p1 != nil {
        cur.Next = p1
    } else if p2 != nil {
        cur.Next = p2
    }

    return head.Next
}
```