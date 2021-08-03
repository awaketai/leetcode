package main

import (
	"fmt"
	"math"
)

// 695. Max Area of Island

// You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
//
//The area of an island is the number of cells with a value 1 in the island.
//
//Return the maximum area of an island in grid. If there is no island, return 0.

var island1 = [][]int{
	{1,1,1,1,0},
	{1,1,0,1,0},
	{1,1,0,0,0},
	{0,0,0,0,0},
}
var island2 = [][]int{
	{1,1,0,0,0},
	{1,1,0,0,0},
	{0,0,1,0,0},
	{0,0,0,1,1},
}

var island3 = [][]int{
	{0,0,1,0,0,0,0,1,0,0,0,0,0},
	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	{0,1,1,0,1,0,0,0,0,0,0,0,0},
	{0,1,0,0,1,1,0,0,1,0,1,0,0},
	{0,1,0,0,1,1,0,0,1,1,1,0,0},
	{0,0,0,0,0,0,0,0,0,0,1,0,0},
	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	{0,0,0,0,0,0,0,1,1,0,0,0,0},
}

var island4 = [][]int{
	{0,0,0,0,0,0,0,0},
}

var island5 = [][]int{
	{1,1,0,0,0},
	{1,1,0,0,0},
	{0,0,0,1,1},
	{0,0,0,1,1},
}

func main() {
	//len1 := len(island4)
	//len2 := len(island4[0])
	//fmt.Println(len1,len2)
	res := maxAreaOfIsland(island1)
	fmt.Println(res)
}

// https://leetcode-cn.com/problems/number-of-islands/solution/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/

func maxAreaOfIsland(grid [][]int) int {
	var res = 0
	for r := 0;r < len(grid);r++ {
		for c := 0;c < len(grid[0]);c++ {
			if grid[r][c] == 1 {
				a := area(grid,r,c)
				fmt.Println(a)
				res = int(math.Max(float64(res), float64(a)))
			}
		}
	}
	return res
}

// 求岛屿面积：每遍历到一个格子，就把面积+1
func area(grid [][]int,r,c int) int {
	if !inArea(grid,r,c){
		return 0
	}
	// 如果是海洋(0)或者已被标记为循环过(2)
	// 则不再循环
	if grid[r][c] != 1 {
		return 0
	}
	// 将已遍历过的岛屿标记为2，避免重复遍历
	grid[r][c] = 2
	// 每遍历到一个格子，如果是岛屿，在将面积+1
	return 1 + area(grid,r-1,c) + area(grid,r+1,c) + area(grid,r,c-1) + area(grid,r,c+1)
}

// 判断坐标(r,c)是否在网格中
func inArea(grid [][]int,r,c int) bool {
	// len(grid):横坐标长度，行
	// grid[0]：竖坐标长度，列
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
