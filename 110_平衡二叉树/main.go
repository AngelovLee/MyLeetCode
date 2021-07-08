// 给定一个二叉树，判断它是否是高度平衡的二叉树。
// 本题中，一棵高度平衡二叉树定义为：
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
package main

import "math"

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

func isBalanced(node *TreeNode) bool {

	if node == nil {
		return true
	}
	//两边最大深度差  大于2
	return isBalanced(node.Left) && isBalanced(node.Right) && // 左右子树是否是平衡树
		math.Abs(height(node.Left)-height(node.Right)) < 2 // 当前节点是否是平衡树

}

//计算节点最大深度
func height(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	//     左右子树深度的最大深度 + 当前节点深度 1
	return math.Max(height(node.Left), height(node.Right)) + 1
}
