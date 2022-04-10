package main

/*
	从第一个元素开始遍历， 到第二个元素s1，思考第二个元素s2 加到第一元素上构成的子序列是增益么
	1. s1是增益； 说明 s1+s2 构成当前最大子序列, 因此将s1 加至 s2 继续前行
	2. s2是增益： 说明s2 单独构成最大子序列
	https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/
	f(i) = max{ f(i−1)+nums[i] , nums[i] }
*/

func maxSubArray(nums []int) int {
	max := nums[0]
	// 边界处理 因为数组从0开始 所以数组边界 小于 长度
	for i := 1; i < len(nums); i++ {
		// 之前的是增益 则
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
