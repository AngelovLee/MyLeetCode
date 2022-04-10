package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的后序遍历：按照访问左子树——右子树——根节点

func postorderTraversal(root *TreeNode) (res []int) {
	var postorderF func(root *TreeNode)
	postorderF = func(root *TreeNode) {
		if root == nil {
			return
		}
		postorderF(root.Left)
		postorderF(root.Right)
		res = append(res, root.Val)
	}
	postorderF(root)
	return res
}
