// 给定一个无重复元素的数组 candidates 和一个目标数 target ，
// 找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
//
//说明：
//	所有数字（包括 target）都是正整数。
//	解集不能包含重复的组合。         (此题难点)

package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

// 二叉树 分支的条件是选择或者不选择
func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{} // 最好不用，make 方法构造
	// 为什么搞成闭包， 因为要从外函数引用 candidates，避免传来传去
	//
	var drillDown func(target, level int, comb []int)
	drillDown = func(target, level int, comb []int) {
		if level == len(candidates) {
			return
		}
		if target == 0 {
			// 对切片深拷贝 append([]int(nil), comb...)
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过 该数 后面都不选择该数
		drillDown(target, level+1, comb)
		// 选择当前数 level,
		if target-candidates[level] >= 0 {
			// 注意不是 level + 1 因为每个数字可以被无限制重复选取，因此搜索的下标仍为 level
			drillDown(target-candidates[level], level, append(comb, candidates[level]))
		}
	}

	// 从下标 0 开始
	drillDown(target, 0, []int{})

	return ans
}

// var res = make([][]int, 0) // 不能写成 make([][]int, 1)
// func combinationSum(candidates []int, target int) [][]int {
//     drillDown(candidates, target, []int{}, 0)
//     return res
// }
// // 一定明确 selectIdx 是指当前将要选择的下标
// func drillDown(candidates []int, target int, selected []int, selectIdx int) {

//     // 当前下标已经越界  直接返回
//     if selectIdx == len(candidates) {
//         return
//     }
//     if target == 0 {
//         res = append(res,  append([]int(nil), selected...))
//         return
//     }

//     // 放弃当前下标
//     drillDown(candidates, target, selected, selectIdx+1)
//     // 选择当前下标
//     if target-candidates[selectIdx] >= 0 {
//         drillDown(candidates, target-candidates[selectIdx], append(selected, candidates[selectIdx]), selectIdx)
//     }

// }
