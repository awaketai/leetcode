package main

import (
	"fmt"
)

// 剑指 Offer 40. 最小的k个数

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

var a1 = []int{4,5,1,6,2,7,3,8}
func main()  {
	ret := getLeastNumbers(a1,4)
	fmt.Println(ret)
}
// 1.使用归并排序，对数列进行排序
// 2.返回指定的元素

// 插入排序：
// 1.创建有序序列，此处视数列第一个元素为有序数列
// 2.扫描剩余所有未排序数列，将每个元素放到有序数列合适位置

// 归并排序
// 1.申请空间存放合并后的数据
// 2.将数据列分割为两两一组
// 3.继续对数据列进行分割，直到每组都只有一个元素
// 4.有序合并分割之后的数据列
func getLeastNumbers(arr []int, k int) []int {
	// 1.使用归并排序，对数据列进行排序
	// 2.从排序的结果中返回元素
	ret := make([]int,0)
	length := len(arr)
	if length == 0 {
		return ret
	}
	// 先排序，然后返回指定数量值
	res := QuickSort(arr)
	return res[:k]
}

// 归并排序
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	// 分割数据列
	left := arr[:middle]
	right := arr[middle:]
	// 不断的对数据列进行分割，直到每个组都只有一个元素,即length<2
	return merge(MergeSort(left),MergeSort(right))
}

func merge(left,right []int) []int {
	var res []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] >= right[0] {
			res = append(res,right[0])
			// 右侧数列指针右移一位
			right = right[1:]
		}else{
			res = append(res,left[0])
			// 左侧指针左移一位
			left = left[1:]
		}
	}
	// 将剩余的有序数列合并到尾部
	for len(left) != 0 {
		res = append(res,left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		res = append(res,right[0])
		right = right[1:]
	}
	return res
}

// 快速排序
// 1.定基准数(选取中间任意数)
// 2.将数据列中小于基准数的元素放到基准数之前，反之放在之后
// 3.对基准数左边和右边的数据集进行递归排序
func QuickSort(arr []int) []int {

	return _quickSort(arr,0,len(arr) - 1)
}

func _quickSort(arr []int,left,right int) []int {
	if left < right {
		index := partition(arr,left,right)
		_quickSort(arr,left,index - 1)
		_quickSort(arr,index+1,right)
	}
	return arr
}

func partition(arr []int,left,right int) int {
	pivot := left // 基准数
	index := pivot + 1
	for i := index;i <= right;i++ {
		if arr[i] < arr[pivot] {
			swap(arr,i,index)
			index++
		}
	}
	swap(arr,pivot,index-1)
	return index-1
}

func swap(arr []int,i,j int)  {

	arr[i],arr[j] = arr[j],arr[i]
}