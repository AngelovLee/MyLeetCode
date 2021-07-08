package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	// 下面 1 表示当前本身， 是需要加一的
	// 左右子树都存在
	if root.Left != nil && root.Right != nil {
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
		// 仅存在左树
	} else if root.Left != nil {
		return 1 + minDepth(root.Left)
		// 仅存在右树
	} else if root.Right != nil {
		return 1 + minDepth(root.Right)
		// 只有当前
	} else {
		return 1
	}

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
