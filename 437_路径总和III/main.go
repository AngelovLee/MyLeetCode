package main


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
* Definition for a binary tree node.
*
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func pathSum(root *TreeNode, sum int) int {
	// 递归解法
	if root == nil {
		return 0
	}
	return allPath(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func allPath(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cnt := 0
	// 当前节点本身要加 1
	currNodeVal := sum - root.Val
	if currNodeVal == 0 {
		cnt++
	}
	// 左右子节点 以及其子子孙孙节点 可能会有路径和 有就上报
	cnt += allPath(root.Left, currNodeVal)
	cnt += allPath(root.Right, currNodeVal)
	return cnt
}

// 错误解法 存在节点重复处理 根节点仅需处理一次其余节点皆处理下一节点
// func pathSum(root *TreeNode, targetSum int) int {
//     if root == nil {
//         return 0
//     }
//     var total = 0
//     var dfs func(tempRoot *TreeNode, tempTargetSum int)
//     dfs = func(tempRoot *TreeNode, tempTargetSum int) {
//         if tempRoot == nil {
//             return
//         }
//         tempTargetSum -= tempRoot.Val
//         if tempTargetSum == 0 {
//             total++
//             return
//         }
//         dfs(tempRoot.Left, tempTargetSum)
//         dfs(tempRoot.Right, tempTargetSum)
//         dfs(tempRoot.Left, targetSum)
//         dfs(tempRoot.Right, targetSum)
//     }
//     // 函数里处理当前节点
//     dfs(root, targetSum)
//     return total
// }
