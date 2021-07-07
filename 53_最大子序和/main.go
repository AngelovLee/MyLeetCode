package main


func maxSubArray(nums []int) int {
	var ans = nums[0]
	var sum = 0
	for i := 0; i < len(nums); i++ {
		// 正常情况下 前 n 的 sum   是 sum_（n-1）+ i_（n）
		// 前 n-1 的正增益才是需要的，不然还不如不要, 直接拿当前的值作为前 n 的 sum
		sum = nums[i] + max(sum, 0)
		// 每次步进算 更新最大和
		ans = max(ans, sum)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

