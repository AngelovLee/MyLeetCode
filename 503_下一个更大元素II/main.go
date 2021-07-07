//给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。
//数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，
//这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

//链接：https://leetcode-cn.com/problems/next-greater-element-ii/solution/xia-yi-ge-geng-da-yuan-su-ii-by-leetcode-bwam/

package main

import "fmt"

func main() {
	n := 10
	for i := 0; i < n*2-1; i++ {
		fmt.Println(i%n)
	}
}

// 总结： 把这个循环数组「拉直」，即复制该序列的前 n-1n−1 个元素拼接在原序列的后面。这样我们就可以将这个新序列当作普通序列
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	stack := []int{}

	// 通过 n*2-1  和 i%n 定位下标会达到循环遍历的效果
	for i := 0; i < n*2-1; i++ {
		// 主要负责找更大的数，若当前下标 i%n 在本次轮没找到，先入栈顶，会重新循环前 i%n
		// 当前下标的数 大于 当栈顶所存小标对应的数的时候
		// 将该栈顶所存的下标 填入对应下标的结果集 并且栈顶出栈
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return ans
}


