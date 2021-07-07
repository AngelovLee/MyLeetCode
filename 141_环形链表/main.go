package main

func main() {}

type ListNode struct {
	 Val int
	 Next *ListNode
}

// 总结： 1. 空节点与一个节点不成环
func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	for slow, fast := head, head.Next;fast != slow; {

		if fast == nil || fast.Next == nil {
			return false
		}

		fast, slow = fast.Next.Next, slow.Next
	}
	return true

}
