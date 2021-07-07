package main

import "fmt"

func main() {
	fmt.Println(
	search([]int{-1,0,3,5,9,12}, 9))
}

//
// for 循环
// 总结：1.变更起始边界 当 target 小于中间时 以 中间为 end
//		2. 中间, 其中 中间位置的计算 为 start + （end-start）%2
//      3. 右移一位 是除以 2
//      4. 结尾是以 长度-1 为下标，长度作为下标 是要减1 的， 数组和列表小标为 0 是第一位
//      5. 什么时候结束， 从趋势上看 end 是 下降的 start是上升的 当 start > end 的时候 就结束了
func search(sortedNums []int, target int) int {

	start, end := 0, len(sortedNums) - 1
	for start <= end {
		mid := start + (end - start) >> 1
		if sortedNums[mid] == target {
			return mid
		} else if target < sortedNums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

// 递归
func search2(sortedNums []int, target int) int {
	start, end := 0, len(sortedNums) - 1
	return binarySearch(sortedNums, start, end, target)
}
func binarySearch(sortedNums []int, start int, end int, target int) int {

	if start > end {
		return -1
	}

	mid := start + (end - start) >> 1
	if sortedNums[mid] == target {
		return mid
	}

	if target < sortedNums[mid] {
		return binarySearch(sortedNums, start, mid - 1, target)
	} else {
		return binarySearch(sortedNums, mid + 1, end, target)
	}
}