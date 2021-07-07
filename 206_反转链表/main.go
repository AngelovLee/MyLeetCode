package main
// 206. 反转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
// https://leetcode-cn.com/problems/reverse-linked-list/
func main() {}
type ListNode struct {

	Val int
	Next *ListNode
}
// 总结： 返回的是 pre 节点
func reverseList(head *ListNode) *ListNode {
	curr := head
	var pre *ListNode
	for curr != nil {
		// 暂存后继
		next := curr.Next
		curr.Next = pre
		// 移动游标
		pre = curr
		curr = next
	}
	return pre
}
