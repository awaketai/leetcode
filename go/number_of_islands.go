package main

import (
	"fmt"
)

// 200. Number of Islands

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
//
//An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.


var i1 = [][]byte{
	{'1','1','1','1','0'},
	{'1','1','0','1','0'},
	{'1','1','0','0','0'},
	{'0','0','0','0','0'},
}

var i2 = [][]byte{
	{'1','1','0','0','0'},
	{'1','1','0','0','0'},
	{'0','0','1','0','0'},
	{'0','0','0','1','1'},
}
func main()  {
	ret := numIslands(i2)
	fmt.Println(ret)
}

// 按照岛屿进行搜索，搜索到的岛屿标记为0
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	index := 0
	lenR,lenC := len(grid),len(grid[0])
	for r := 0;r < lenR;r++ {
		for c := 0;c < lenC;c++ {
			if grid[r][c] == '1' {
				dfs1(&grid,r,c)
				index++
			}
		}
	}
	return index
}

func dfs1(grid *[][]byte,r,c int) int {
	if !inArea3(*grid,r,c) {
		return 0
	}
	if (*grid)[r][c] != '1' {
		return 0
	}
	(*grid)[r][c] = '0'
	dfs1(grid,r-1,c)
	dfs1(grid,r+1,c)
	dfs1(grid,r,c-1)
	dfs1(grid,r,c+1)
	return 1
}


// 判断坐标(r,c)是否在网格中
func inArea3(grid [][]byte,r,c int) bool {
	// len(grid):横坐标长度，行
	// grid[0]：竖坐标长度，列
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}