package main

/**
* Definition for a binary tree node.
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//  var res [][]int

// func pathSum(root *TreeNode, sum int) [][]int {
// 	res = [][]int{}
// 	dfs(root, sum, []int{})
// 	return res
// }

// func dfs(root *TreeNode, sum int, stack []int) {
// 	if root == nil {
// 		return
// 	}
// 	stack = append(stack, root.Val)
// 	if root.Left == nil && root.Right == nil {
// 		if sum == root.Val {
// 			r := make([]int, len(stack))
// 			copy(r, stack)
// 			res = append(res, r)
// 		}
// 	}
// 	dfs(root.Left, sum-root.Val, stack)
// 	dfs(root.Right, sum-root.Val, stack)
// }

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	var dfs func(root *TreeNode, targetSum int, nodes []int)
	dfs = func(root *TreeNode, targetSum int, nodes []int) {
		if root == nil {
			return
		}
		// 错误做法 因为 nodes 切片， 切片是数组的引用，append 的切片，在后续引用变更中会被改变
		nodes = append(nodes, root.Val)
		if root.Left == nil && root.Right == nil && targetSum == root.Val {
			ans = append(ans, nodes)
		}
		// 正确做法1 记得深拷贝切片
		// nodes = append(nodes, root.Val)
		// if root.Left == nil && root.Right == nil && targetSum == root.Val {
		//     r := make([]int, len(nodes))
		//     copy(r, nodes)
		//     ans = append(ans, r)

		// }

		// 正确做法2 记得深拷贝切片
		dfs(root.Right, targetSum - root.Val, append([]int(nil), nodes...))
		dfs(root.Left, targetSum - root.Val, append([]int(nil), nodes...))

	}
	dfs(root, targetSum, []int{})
	return ans
}
