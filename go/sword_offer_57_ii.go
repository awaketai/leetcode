package main

import "fmt"

// 剑指 Offer 57 - II. 和为s的连续正数序列

// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//
//序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

func main()  {
	ret := findContinuousSequence(15)
	fmt.Println(ret)
}

// 滑动窗口:窗口的左边界和右边界只能向右移动
// 需要关注的是窗口中所有数字的和，当窗口向右移动时，要加上右边的值
// 当窗口缩小时，需要减去左边的值
// 当窗口内所有数字和小于target时，窗口需要扩大，大于target时，窗口需要缩小
//
func findContinuousSequence(target int) [][]int {
	left,right,sum := 1,2,3 // 滑动窗口左右边界,窗口中数字和
	res := make([][]int,0)
	for left < right {
		if sum == target {
			cur := make([]int,0)
			for i := left;i <= right;i++ {
				cur = append(cur,i)
			}
			res = append(res,cur)
			sum -= left
			left++
		}else if sum > target {
			// 如果和大于target，则缩小窗口，做边界右移
			sum -= left
			left++
		}else{
			// 如果和小于target，则窗口扩大，右边界右移
			right++
			sum += right
		}
	}
	return res
}