package main

import "fmt"

// 463. Island Perimeter

// You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.
//
//Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).
//
//The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island. One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.


var p1 = [][]int{
	{0,1,0,0},
	{1,1,1,0},
	{0,1,0,0},
	{1,1,0,0},
}
var p2 = [][]int{
	{1},
}
var p3 = [][]int{
	{1,0},
}

func main()  {
	ret := islandPerimeter(p3)
	fmt.Println(ret)
}

// 思路来自leetcode-cn
// https://leetcode-cn.com/problems/number-of-islands/solution/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/
func islandPerimeter(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	lenR,lenC := len(grid),len(grid[0])
	for r := 0;r < lenR;r++ {
		for c := 0;c <lenC;c++ {
			if grid[r][c] == 1 {
				return dfs(grid,r,c)
			}
		}
	}
	return 0
}

func dfs(grid [][]int,r,c int) int {
	// 如果超出网格范围，则对应一条边[黄色边]
	if !inArea2(grid,r,c){
		return 1
	}
	// 如果当前遍历到海洋格子，则对应一条边[蓝色边]
	if grid[r][c] == 0 {
		return 1
	}
	// 如果是陆地格子
	if grid[r][c] != 1 {
		return 0
	}
	// 将已遍历过的格子标记为2，避免重复遍历
	grid[r][c] = 2
	return dfs(grid,r,c-1) + dfs(grid,r,c+1) + dfs(grid,r-1,c) + dfs(grid,r+1,c)
}

// 判断坐标(r,c)是否在网格中
func inArea2(grid [][]int,r,c int) bool {
	// len(grid):横坐标长度，行
	// grid[0]：竖坐标长度，列
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
