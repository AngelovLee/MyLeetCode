package main

// 3. 无重复字符的最长子串
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

// 双指针法
// 总结： 1. 定义左边界和右边界
//       2. 左边界移动到已存在字符的下标 + 1 的位置
//       3. 右边每次都 + 1
//       4. 循环结束条件  r < cnt;  cnt 从 1 开始，左右边界指针从 0 开始
//       5. 注意 cnt == 是否可以包含进循环内
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	for cnt, l, r, existedMap := len(s), 0, 0, make(map[uint8]int);r < cnt;r++ {

		if existedIdx, ok := existedMap[s[r]];ok && existedIdx >= l {
			l = existedIdx + 1
		}

		if r - l + 1 > maxLen {
			maxLen = r - l + 1
		}

		existedMap[s[r]] = r

	}

	return maxLen
}

