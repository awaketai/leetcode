package main

import (
	"fmt"
)

//剑指 Offer 04. 二维数组中的查找

// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//

var f1 = [][]int{

	{1,   4,  7, 11, 15},
	{2,   5,  8, 12, 19},
	{3,   6,  9, 16, 22},
	{10, 13, 14, 17, 24},
	{18, 21, 23, 26, 30},
}

func main()  {
	ret := findNumberIn2DArray2(f1,20)
	fmt.Println(ret)
}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	ret := binarySearch1(matrix,target)
	fmt.Println(ret)
	return ret
}

// 按照行二分查找
func binarySearch1(matrix [][]int,target int) bool {
	if len(matrix) == 0 {
		return false
	}
	// 对每一行按照二分法进行查找
	lenRow := len(matrix)
	for i := 0;i < lenRow;i++ {
		start := 0
		end := len(matrix[0]) - 1
		// a[0][0]
		for start <= end {
			mid := start + ((end - start) / 2)
			if matrix[i][mid] == target {
				return true
			}
			if matrix[i][mid] > target {
				// 在左边找
				end = mid - 1
			}else{
				// 在右边找
				start = mid + 1
			}
		}
	}
	return false
}

// 以二维数组左下角为原点，建立直角坐标
// 1.当前数字大于了查找数，查找上移
// 2.当前数字小于查找数，坐标右移
func findNumberIn2DArray2(matrix [][]int,target int) bool {
	if len(matrix) == 0 {
		return false
	}
	lenRow := len(matrix)
	lenColumn := len(matrix[0])
	x,y := lenRow-1,0 // 从二维数组左下角为坐标原点
	for x >= 0 && y < lenColumn {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] < target {
			// 当前元素 < 目标元素，从右侧找
			y++
		}else{
			// 当前元素> 目标元素，从上侧找
			x--
		}
	}
	return false
}