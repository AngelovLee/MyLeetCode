package main

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var backTrack func(nums []int, target, cur int)
	backTrack = func(nums []int, target, cur int) {
		if len(nums) == k && target == 0 {
			res = append(res, nums)
			return
		}
		if target < 0 {
			return
		}
		for i := cur; i < 10; i++ {
			// 选该数
			nums = append(nums, i)
			backTrack(nums, target - i, i+1)
			// 回退该数，为下一个留空间，因为这一层只
			nums = append([]int{}, nums[:len(nums)-1]...)
		}
	}
	backTrack([]int{}, n, 1)
	return res
}
