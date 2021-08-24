package main

// 剑指 Offer 16. 数值的整数次方

// 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

// 假设十进制数 5,其二进制为0101 -> 2^2 + 2^0
// 2 ^ 5 即 2 ^ (2^2 + 2^0) -> 2 ^ (2^2) * 2^(2^0)
// 待理解...
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	b := n
	res := 1.0
	// 如果指数为负
	// 2 ^ -5 = 1 / 2^5
	if b < 0 {
		x = 1 / x
		b = -b
	}
	for b > 0 {
		// 如果最右是1
		if (b & 1) == 1 {
			res *= x
		}
		x *= x
		b >>= 1
	}
	return res
}