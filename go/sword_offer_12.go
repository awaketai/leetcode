package main

import (
	"fmt"
)

//剑指 Offer 12. 矩阵中的路径

//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。


var board1 = [][]byte{
	{'A','B','C','E'},
	{'S','F','C','S'},
	{'A','D','E','E'},
}
var word1 = "ABCCED"

var board2 = [][]byte{
	{'a','b'},
	{'c','d'},
}

var word2 = "abcd"

func main() {

	fmt.Println(string(word1[2]))
}

// 未完待解...
// 从左上角开始查找，如果四周相邻的元素有相等的，则继续找
// 否则停止

func exist(board [][]byte, word string) bool {
	lenR,lenC := len(board),len(board[0])
	if lenR == 0 {
		return false
	}
	for r := 0;r < lenR;r++ {
		for c := 0;c < lenC;c++ {

		}
	}


	return true
}

// 退出条件：1.越界
// 2.
func dfs4(board [][]byte,r,c,k int,word string) bool {
	if !inArea4(board,r,c){
		return false
	}
	if board[r][c] != word[k] {
		return false
	}
	if board[r][c] == ' ' {
		return false
	}
	// 将已访问过的元素修改，避免重复访问
	board[r][c] = ' '
	dfs4(board,r,c-1,k + 1,word)
	dfs4(board,r,c+1,k + 1,word)
	dfs4(board,r-1,c,k + 1,word)
	dfs4(board,r+1,c,k + 1,word)
	return true
}

func inArea4(board [][]byte,r,c int) bool {

	return 0 <= r && r < len(board) && 0 <= c && c < len(board[0])
}
