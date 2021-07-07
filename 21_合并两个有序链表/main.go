package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 总结： 合并链表 不可在愿链表上做变更 必须另起一个头节点
//        排版要工整对其， 方便排错
func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	res := &ListNode{}
	if l1.Val < l2.Val {
		res = l1
		res.Next = mergeTwoList(l1.Next, l2)
	} else {
		res = l2
		res.Next = mergeTwoList(l1, l2.Next)
	}
	return res
}


