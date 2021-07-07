package main

func main() {

}

type ListNode struct {
	Val int
	Next *ListNode
}
// 总结： go中 append 与顺序无关
// 方法1
func reversePrint(head *ListNode) []int {
	count := 0
	for p:= head;p != nil;p = p.Next{
		count++
	}
	res := make([]int,count)
	for ;head!= nil;head=head.Next{
		res[count-1] = head.Val
		count--
	}
	return res
}

// 方法2
//func reversePrint(head *ListNode) ([]int) {
//	if head == nil {return nil}
//	return append(reversePrint(head.Next), head.Val)
//}

