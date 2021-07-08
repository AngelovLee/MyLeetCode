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

// 总结： 前序遍历，遍历时记录层级，记录层级的值
func levelOrder(root *TreeNode) (res [][]int) {
	var preorder func(node *TreeNode, depath int)
	preorder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], node.Val)
		preorder(node.Left, depth+1)
		preorder(node.Right, depth+1)
		return
	}
	preorder(root, 0)
	return
}

// 方法二
//func levelOrder(root *TreeNode) [][]int {
//	ret := [][]int{}
//	if root == nil {
//		return ret
//	}
//	q := []*TreeNode{root}
//	for i := 0; len(q) > 0; i++ {
//		ret = append(ret, []int{})
//		p := []*TreeNode{}
//		for j := 0; j < len(q); j++ {
//			node := q[j]
//			ret[i] = append(ret[i], node.Val)
//			if node.Left != nil {
//				p = append(p, node.Left)
//			}
//			if node.Right != nil {
//				p = append(p, node.Right)
//			}
//		}
//		q = p
//	}
//	return ret
//}
