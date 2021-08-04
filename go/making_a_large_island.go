package main

import (
	"fmt"
	"math"
)

/// 827. Making A Large Island

// You are given an n x n binary matrix grid. You are allowed to change at most one 0 to be 1.
//
//Return the size of the largest island in grid after applying this operation.
//
//An island is a 4-directionally connected group of 1s.

var island32 = [][]int{
	{0,0,1,0,0,0,0,1,0,0,0,0,0},
	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	{0,1,1,0,1,0,0,0,0,0,0,0,0},
	{0,1,0,0,1,1,0,0,1,0,1,0,0},
	{0,1,0,0,1,1,0,0,1,1,1,0,0},
	{0,0,0,0,0,0,0,0,0,0,1,0,0},
	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	{0,0,0,0,0,0,0,1,1,0,0,0,0},
}

var island33 = [][]int{
	{1, 0},
	{0, 1},
}

var island34 = [][]int {
	{1, 1}, {1, 0},
}

var island35 = [][]int{
	{1, 1},
	{1, 1},
}


func main()  {
	ret := largestIsland(island35)
	fmt.Println(ret)
}
// 思路来自leetcode-cn
// 1.先标记岛屿每块岛屿标记为一个索引
// 2.单独记录每个岛屿的面积
// 3.遍历海洋，查看那个海洋块相邻的两个岛屿面积最大
func largestIsland(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 1
	}
	lenR,lenC := len(grid),len(grid[0])
	index := 1 // 岛屿索引标记起始索引
	area := make(map[int]int) // 记录指定编号岛屿面积 2:3
	maxArea := 0
	for r := 0;r < lenR;r ++ {
		for c := 0;c < lenC;c++ {
			// 遍历岛屿，并标记索引
			if grid[r][c] == 1 {
				index += 1
				curArea := island(grid,r,c,index)
				// 当前索引岛屿面积保存
				area[index] = curArea
				maxArea = int(math.Max(float64(maxArea), float64(curArea)))
			}

		}
	}
	// 如果没有陆地，则造一块
	if maxArea == 0 {
		return 1
	}
	for r := 0;r < lenR;r ++ {
		for c := 0;c < lenC;c++ {
			// 遍历海洋，查找海洋格子相邻的岛屿面积最大的
			if grid[r][c] == 0 {
				oceanNeighbour := ocean(grid,r,c)
				if len(oceanNeighbour) == 0 {
					continue
				}
				// 当前海洋四周岛屿以岛屿编号：1 的方式记录到了map中
				// 将四周岛屿面积想加，并和岛屿最大面积比较
				islandMaxArea := 1
				for i,_ := range oceanNeighbour {
					islandMaxArea += area[i]
				}
				maxArea = int(math.Max(float64(maxArea),float64(islandMaxArea)))
			}
		}
	}
	return maxArea
}

// 海洋中查找四周岛屿，并将海洋周围岛屿以岛屿编号：1的方式记录到map中
func ocean(grid [][]int,r,c int) map[int]int {
	maxArea := make(map[int]int)
	// 如果当前坐标超出网格
	if inArea1(grid,r,c-1) && grid[r][c-1] != 0 {
		maxArea[grid[r][c-1]] = 1
	}
	if inArea1(grid,r,c+1) && grid[r][c+1] != 0 {
		maxArea[grid[r][c+1]] = 1
	}
	if inArea1(grid,r-1,c) && grid[r-1][c] != 0 {
		maxArea[grid[r-1][c]] = 1
	}
	if inArea1(grid,r+1,c) && grid[r+1][c] != 0 {
		maxArea[grid[r+1][c]] = 1
	}
	return maxArea
}

func island(grid [][]int,r,c,index int) int {
	// 如果当前坐标超出网格
	if !inArea1(grid,r,c){
		return 0
	}
	// 遍历所有岛屿并标记岛屿索引
	if grid[r][c] != 1 {
		return 0
	}
	// 标记岛屿为指定索引
	grid[r][c] = index
	return 1 + island(grid,r-1,c,index) +
	island(grid,r,c-1,index) +
	island(grid,r+1,c,index) +
	island(grid,r,c+1,index)
}

// (r,c)坐标的元素是否在网格之中
func inArea1(grid [][]int,r,c int) bool {

	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
