package main

// 	剑指 Offer 22. 链表中倒数第k个节点
//	输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//
//	例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
//	链接：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof
func main() {}

type ListNode struct {
	Val int
	Next *ListNode
}

// 总结： 1. 数组、切片、字符串、整数 都是线性对称结构
//       2. 凡是都要考虑边界（越界问题）
func reverse_last_k(head *ListNode, k int) *ListNode{

	count := k
	curr := head
	for curr != nil {
		curr = curr.Next
		count--
	}

	// 处理边界问题
	if count > 0 {
		return nil
	}
	if count == 0 {
		return head
	}

	curr = head
	for count < 0 {
		curr = curr.Next
		count ++
	}
	return curr
}
