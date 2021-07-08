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

func levelOrderBottom(root *TreeNode) [][]int {
	//深度优先遍历(dfs)
	result := make([][]int, 0)
	var preorder func(node *TreeNode, depath int)
	preorder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(result) {
			result = append(result, []int{})
		}
		result[depth] = append(result[depth], node.Val)
		preorder(node.Left, depth+1)
		preorder(node.Right, depth+1)
		return
	}
	preorder(root, 0)

	//数组翻转
	resultLength := len(result)
	left := 0
	right := resultLength - 1
	for left < right {
		temp := result[left]
		result[left] = result[right]
		result[right] = temp

		left++
		right--
	}

	return result
}
