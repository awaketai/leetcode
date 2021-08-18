package main

import (
	"fmt"
)

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面

// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

func main()  {
	var n = []int{1,2,3,4,5,6}
	ret := exchange2(n)
	fmt.Println(ret)
}

// 方法一：双指针法
// 首位双指针left,right
// left往右移，直到值为偶数
// right往左移，直到值为奇数
// 交换left和right的值
// 直到left和right相遇
func exchange(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nums
	}
	left,right := 0,length - 1
	for left < right {
		if (nums[left] & 1) != 0 {
			// 左指针当前位置是奇数，不做交换，左指针前移
			left++
			continue
		}
		if (nums[right] & 1) != 1 {
			// 右指针当前位置是偶数，不做交换，右指针后移
			right--
			continue
		}
		// 如果左边是偶数，右边是奇数，则交换
		nums[left],nums[right] = nums[right],nums[left]
	}
	return nums
}

// 方法二：快慢指针法
// 定义快慢指针，fast和low，fast在前，low在后
// fast向前搜索奇数位置，low指向下一个奇数应当存储的位置
// fast向前移动，当搜索到奇数时，将ta和nums[low]交换，此时low向前移动一个位置
// 重复上述操作，直到fast指向数组末尾
func exchange2(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nums
	}

	fast,low := 0,0
	for fast < length {
		if (nums[fast] & 1) == 1 {
			// 如果fast当前位置是偶数，z则交换
			nums[fast],nums[low] = nums[low],nums[fast]
			// 慢指针移动到下一个要交换的位置
			low++
		}
		fast++
	}
	return nums
}