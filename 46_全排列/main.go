// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

// 解题思路
// 如果数组只有一个数 a， 那全排列就是它本身 a，
// 如果有两个数 a, b。 把它拆分开， 把a 追加到 b 的全排列 就是 [b, a]。 把b追加到a的全排列就是 [a, b]
// 好，那么三个数 a,b,c。
// 把 a 拿出来 追加到 b，c的全排列得到 [c,b,a] 和 [b,c,a]。
// 再把b拿出来 追加到 a, c的全排列得到 [c,a,b]和[a,c,b]。
// 再把c拿出来 追加到 a, b的全排列得到 [b,a,c]和[a,b,c]。
// 依次类推。。。

// 所以递推公式 就是数组中每次拿出一个，把它追加到剩余数组的全排列中
// 递归的出口是，数组只有一个数时，把自己返回
package main

import "fmt"

func main() {
	selected := []int{1, 2, 3}
	fmt.Println(append(append([]int{}, selected...), 1))
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := [][]int{}

	for i, num := range nums {
		// 把num从 nums 拿出去 得到tmp
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])

		// sub 是把num 拿出去后，数组中剩余数据的全排列
		sub := permute(tmp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}
	}
	return res
}
