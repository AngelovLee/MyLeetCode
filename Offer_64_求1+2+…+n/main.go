package main

// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

// 总结： 自相似 -> 递归
func sumNums(n int) int {
	if n == 0 {
		return n
	}
	return sumNums(n-1)+n
}
