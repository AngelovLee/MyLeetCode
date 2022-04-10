package main

import "math"

/*
   直到到第i天买通行证所花的所有钱
   dp[i] = costs[0] + dp[i+1]
   dp[i] = costs[1] + dp[i+7]
   dp[i] = costs[2] + dp[i+30]

   dp[i] = min(costs[0] + dp[i+1], costs[1] + dp[i+7], costs[2] + dp[i+30])
*/
func mincostTickets(days []int, costs []int) int {
	dp := make([]int, 396)
	dayMap := map[int]bool{}
	for _, day := range days {
		dayMap[day] = true
	}
	for day := 365; day > 0; day-- {
		if dayMap[day] {
			dp[day] = int(math.Min(math.Min(float64(costs[0]+dp[day+1]), float64(costs[1]+dp[day+7])), float64(costs[2]+dp[day+30])))
		} else {
			dp[day] = dp[day+1]
		}
	}

	return dp[1]
}
