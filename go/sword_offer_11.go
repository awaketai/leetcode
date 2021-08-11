package main

import "fmt"

// 剑指 Offer 11. 旋转数组的最小数字


//把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。  
//

var a11 = []int{3,4,5,1,2}
var a22 = []int{2,2,2,0,1}
var a33 = []int{7,0,1,1,1,1,1,2,3,4}
var a44 = []int{3,4,5,6,1,2,3}

func main()  {
	ret := minArray(a44)
	fmt.Println(ret)
}
// 使用二分法
// [3,4,5,1,2]
// [2,2,2,0,1]
// [7,0,1,1,1,1,1,2,3,4]
// 用左边值也mid值进行比较不能很好的减治
func minArray(numbers []int) int {
	length := len(numbers)
	if length == 0 {
		return 0
	}
	end := length - 1
	start := 0
	for start <= end {
		mid := start + (end - start) / 2
		// 如果end位置元素小于中间位置元素，则最小值在右侧
		if numbers[end] < numbers[mid] {
			start = mid + 1
		}else if numbers[end] > numbers[mid] {
			// 如果end位置元素大于中间元素，则最小值在左侧,为了不漏值，不减-1
			// 右侧的值是可以丢弃的
			end = mid
		}else {
			// 如果中间位置元素和最右侧元素相等，则移动右侧指针，一个个找
			end--
		}
	}
	return numbers[start]
}