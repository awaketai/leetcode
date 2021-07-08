package main

import "fmt"

func main()  {
	var arr = make([]int,1000)
	for i := 0;i < 1000;i++ {
		arr[i] = i
	}
	//fmt.Println(arr)
	fmt.Println(binarySearch(arr,1730))
}


func binarySearch(arr []int,target int) int {
	var start = 0 // 查找起点
	var end = len(arr) - 1 // 查找终点
	var mid int // 查找中位数
	for start <= end {
		mid = start + ((end - start ) / 2)

		fmt.Println("mid:",mid)
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			start = mid + 1
		}else{
			end = mid - 1
		}
	}
	return -1
}
