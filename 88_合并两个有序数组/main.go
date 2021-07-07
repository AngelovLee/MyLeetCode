package main
// 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。
//链接：https://leetcode-cn.com/problems/merge-sorted-array


func main() {

}

// m + n 代表游标数
// m+n 循环
// 比大小 大 往num1后填充
// 总结： 1. 既然有循环比较 就有游标来计数
//       2. go 里面拿长度定位里面的元素 是 length - 1; length - 1 >= 0
//       3. 循环（步进）循环的目标是实际的对象还是 对象的组合， 此题是对象的组合（步进时互斥）
func merge(nums1 []int, m int, nums2 []int, n int) []int {

	//
	// 循环条件 双方游标都没结束时
	// 通过 n 结束情况判断（判断短的） 1. num2 已结束  2. num2 未结束
	for p:=m+n; n>0 && m>0; p-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[p-1] = nums1[m-1]
			m--
		} else {
			nums1[p-1] = nums2[n-1]
			n--
		}
	}
	for ;n-1>=0;n-- {
		nums1[n-1] = nums2[n-1]
	}
	return nums1
}