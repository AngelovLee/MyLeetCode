// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

package main

import (
	"fmt"
	"math"
)

func main() {
	// 考虑如下 特殊情况，最低价在最高价后面， 会不会有问题， 不会
	// 因为 整个循环过程中的 最大利润 都是会通过 股票上涨且比较过后才会变更
	fmt.Println(maxProfit([]int{3,4,5,4,3,2,1,0}))
}

func maxProfit(prices []int) int {
	// 还未持有
	held := math.MaxInt64
	maxValue := 0
	for i := 0; i < len(prices); i++ {
		// 先买才能卖 持续预判 直到买到最低价
		// 高价卖 低价买
		//
		if prices[i] < held {
			held = prices[i]
		// 持续预判 知道拿到最高价
		} else if prices[i]-held > maxValue {
			maxValue = prices[i] - held
		}
	}
	return maxValue
}

