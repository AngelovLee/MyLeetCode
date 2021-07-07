package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-123))
	fmt.Println(reverse(123))
}

// 总结：1. 需注意负数 % 10 等 负最后一位，
//			因此循环终结条件为 x!=0
//		2. 结果边界处理
//      3. res := 0 要在 for 循环外部定义
func reverse(x int) int {

	res := 0
	for ;x!=0; x/=10 {
		res *= 10
		res += x%10
		if res < math.MinInt32 || res > math.MaxInt32 {
			return 0
		}
	}
	return res
}