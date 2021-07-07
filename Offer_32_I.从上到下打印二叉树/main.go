package main

type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	treeArr := []*TreeNode{root}
	resultList := make([]int,0)

	for len(treeArr) != 0 {
		currNode := treeArr[0]
		treeArr = treeArr[1:]
		if currNode.Left != nil {
			treeArr = append(treeArr,currNode.Left)
		}

		if currNode.Right != nil {
			treeArr = append(treeArr,currNode.Right)
		}

		resultList = append(resultList,currNode.Val)
	}

	return resultList
}
