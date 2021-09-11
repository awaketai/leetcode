package main

// 剑指 Offer 44. 数字序列中某一位的数字

// 数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
//
//请写一个函数，求任意第n位对应的数字。
//

func main()  {

}
// 思路来自：
// https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/solution/mian-shi-ti-44-shu-zi-xu-lie-zhong-mou-yi-wei-de-6/

// 数字范围  位数  数字数量  数位数量
// 1~9       1      9       9
// 10~99     2      90      180
// 100~999   3      900     2700
// start~end  digit  9*start  9*start*digit
// 位数地推公式：digit=digit+1
// 其实数字递推公式：start=start*10
// 数位数量计算公式：count=9*start*digit

// 解法：
// 1.确定n所在数字的位数，记为digit
// 2.确定n所在的数字，记为num
// 3.确定n是num中的哪一位数，并返回结果

func findNthDigit(n int) int {
	// 1.确定n所在数字位数
	digit,start,count := 1,1,9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = 0 * start * digit
	}
	return 0
}
