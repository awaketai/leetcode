package main

import "fmt"

// 不修改数组找出重复的数字
// 在一个长度为n+1的数组里的所有数字都在1~n的范围内，所以数组中至少有一个数字是重复的。
// 请找出数组中任意一个重复的数字，但不能修改输入的数组。例如，如果输入长度为8的数组{2,3,5,4,3,2,6,7}
// 那么对应的输出是重复的数字2或者3

// 抽屉原理：有十个苹果，将其放入9个抽屉，总有一个抽屉一定是要放两个苹果的

func duplicationInArrayNoEdit(nums []int) int {
	start,end := 1,len(nums) - 1
	for start <= end {
		// 8-1 7 0111 0011 3
		// 4-1 3 0011 0001 1
		mid := ((end - start) >> 1) + start
		// 计算左右两边分别对应的数字个数
		// 假设左边为1-4,计算左边1-4范围内的数字，如果超过四个则左边重复
		// 右边为5-8，计算右边5-8范围内的数字
		count := countRange(nums,len(nums),start,mid)
		if end == start {
			if count > 1 {
				return start
			}else{
				break
			}
		}
		if count > (mid - start + 1){
			end = mid
		}else{
			start = mid + 1
		}

	}
	return -1
}

func countRange(nums []int,n,start,end int) int {
	count := 0
	for i := 0;i < n;i++ {
		if nums[i] >= start && nums[i] <= end {
			count++
		}
	}
	return count
}

func main()  {
	nums1 := []int{2,3,5,4,3,2,6,7}
	ret := duplicationInArrayNoEdit(nums1)
	fmt.Println(ret)
}



