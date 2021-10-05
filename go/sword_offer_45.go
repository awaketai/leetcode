package main

import (
	"log"
	"sort"
	"strconv"
)

// 剑指 Offer 45. 把数组排成最小的数

// 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

// 示例 1:
//
//输入: [10,2]
//输出: "102"

// 自定义排序规则 +:连接符
// 若 x+y > y+x 则 x > y
// 若 x+y < y+x 则 x < y
// 若x < y：排序后x应在y左边，否则反之
// 依据此思路，使用排序算法进行排序

func main()  {
	nums1 := []int{10,2}
	nums2 := []int{3,30,34,5,9}
	nums3 := []int{1,2,3,1}
	str1 := minNumber(nums1)
	if "102" != str1 {
		log.Fatalf("test error: expected %s,actual %s","102",str1)
	}
	str2 := minNumber(nums2)
	if "3033459" != str2 {
		log.Fatalf("test error: expected %s,actual %s","3033459",str2)
	}
	str3 := minNumber(nums3)
	if "1123" != str3 {
		log.Fatalf("test error: expected %s,actual %s","1123",str3)
	}
	log.Println("test success!")
}

func minNumber(nums []int) string {
	// 1.将元素转换为字符串
	var strs = []string{}
	for _,v := range nums {
		strs = append(strs,strconv.Itoa(v))
	}
	// 2.使用快速排序，进行排序 
	_quickSort(strs,0,len(strs) - 1)
	// 使用函数排序
	//sort.Slice(strs, func(i, j int) bool {
	//	return strs[i]+strs[j] <= strs[j]+strs[i]
	//})
	// 3.对排序后的切片连接为字符串
	var str string
	for _,vo := range strs {
		str += vo
	}
	return str
}

// 方法一：自定义快速排序
func _quickSort(nums []string,left,right int)  {
	if left < right {
		pivot := doQuickSort(nums,left,right)
		_quickSort(nums,left,pivot - 1)
		_quickSort(nums,pivot+1,right)
	}
}

func doQuickSort(nums []string,left,right int) int {
	pivot := nums[left]
	for left < right {
		// 将小于基准元素的放左边，反之放右边
		for left < right && pivot + nums[right] <= nums[right]+pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] + pivot <= pivot+nums[left] {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}

