package main

// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//你可以认为每种硬币的数量是无限的。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/coin-change

import "math"

// 总结
//		1. 状态转移方程： F(n) = min(F(n-coin_1)+1，F(n-coin_2)+1 ，..., F(n-coin_i)+1)   (coin_i 为 硬币面额)
//      2. math.MaxInt32 可替换为 amount + 1
func coinChange(coins []int, amount int) int {
	if amount < 1 && len(coins) < 1 {
		return -1
	}
	memo := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		memo[i] = math.MaxInt32
		for _, c := range coins {
			if i >= c && memo[i] > memo[i-c]+1 {
				memo[i] = memo[i-c] + 1
			}
		}
	}
	if memo[amount] == math.MaxInt32 { return -1 }
	return memo[amount]
}

