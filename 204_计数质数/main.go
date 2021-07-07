package main

// 埃氏筛
//初始化长度 O(n)的标记数组，表示这个数组是否为质数。数组初始化所有的数都是质数.
// 从 2 开始将当前数字的倍数全都标记为合数。标记到 n的开方时停止即可

// 每次找当前素数 x 的倍数时，是从 x^2开始的。因为如果 x > 2x>2，那么 2*x2∗x 肯定被素数 2 给过滤了，最小未被过滤的肯定是 x^2
//链接：https://leetcode-cn.com/problems/count-primes/solution/kuai-lai-miao-dong-shai-zhi-shu-by-sweetiee/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func countPrimes(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}
