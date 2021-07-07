package main

import "fmt"

// 队列是fifo 栈是filo
// 要点 从后面添加 从后面拿

type MyQueue struct {
	inStack, outStack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.inStack = append(this.inStack, x)
	fmt.Println(this.inStack)
}

func (this *MyQueue)in2Out() {
	for l:=0; len(this.inStack) > 0;l++ {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}

/** Removes the element from in front of queue and returns that element.
  移出对头
*/
func (this *MyQueue) Pop() int {
	// 为何 有类型这样的判断， 因为保持先进先出
	if len(this.outStack) == 0 {
		this.in2Out()
	}

	res := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return res

}


/** Get the front element.
  从队头拿
*/
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.in2Out()
	}
	fmt.Println("Peek Status:", this.inStack, this.outStack)
	return this.outStack[len(this.outStack)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

func main()  {
	//obj := Constructor()
	//obj.Push(1)
	//obj.Push(2)
	//fmt.Println(obj.Peek())
	//obj.Push(3)
	//obj.Push(4)
	//obj.Push(5)
	//fmt.Println(map[int]int{0:1})
	for k,v :=range map[int]int{0:1, 2:3} {
		fmt.Println(k, v)
	}
}
/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
