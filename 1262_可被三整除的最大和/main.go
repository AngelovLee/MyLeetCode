package main

import "math"

/*
	https://leetcode-cn.com/problems/greatest-sum-divisible-by-three/
	实例：
	输入：nums = [3,6,5,1,8]
	输出：18
	解释：选出数字 3, 6, 1 和 8，它们的和是 18（可被 3 整除的最大和）。

	思路： 从第一个元素e1开始，它与3除取余数有三种情况： 0，1，2， 但每一种情况都有可能在与下一个元素e2相加后在一次，洗牌这三种情况
          所有累加的结果要存，放到下一次判断
*/
func maxSumDivThree(nums []int) int {
	remainders := []int{0, 0, 0}
	currRemainders := []int{0, 0, 0}
	for _, num := range nums {
		for _, r := range remainders {
			currTotal := r + num
			mod := currTotal % 3
			if mod == 0 {
				currRemainders[0] = int(math.Max(float64(currRemainders[0]), float64(currTotal)))
			} else if mod == 1 {
				currRemainders[1] = int(math.Max(float64(currRemainders[1]), float64(currTotal)))
			} else {
				currRemainders[2] = int(math.Max(float64(currRemainders[2]), float64(currTotal)))
			}
		}
		remainders = append([]int{}, currRemainders...)
	}
	return remainders[0]
}
