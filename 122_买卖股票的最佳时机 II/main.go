package main

func maxProfit(prices []int) int {
	n := len(prices)
	// dp0 第一天未买持股票
	// dp1 第一天买了股票
	dp0, dp1 := 0, -prices[0]

	// 从第二天开始
	for i := 1; i < n; i++ {
		// 状态方程如下:
		// 	今天没有持有股票 等于 max(dp0, dp1+prices[i]) 昨天没买股票 昨天买了股票今天卖
		// 	今天持有股票    等于 max(dp1, dp0-prices[i]) 昨天持有股票 昨天没持有股票今天要买
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	// 今天没持有肯定比今天持有钱多，不过这里可以来个 max 比较
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
