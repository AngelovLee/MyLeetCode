//给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
//
//请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
//
//nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
//
//链接：https://leetcode-cn.com/problems/next-greater-element-i

package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	stack := make([]int, 0, len(nums2))
	m := make(map[int]int)
	res := make([]int, len(nums1))

	for i := 0; i < len(nums2); i++ {
		for len(stack) != 0 && stack[len(stack) - 1] < nums2[i] {
			// 记录栈顶元素 的 下一个更大元素是谁
			m[stack[len(stack) - 1]] = nums2[i]
			// 栈顶出栈
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, nums2[i])
	}

	for i := 0; i < len(nums1); i++ {
		if v, ok := m[nums1[i]]; ok {
			res[i] = v
		} else {
			res[i] = -1
		}
	}

	return res
}

