<?php

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
//
//An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.


class Solution {

    /**
     * @param String[][] $grid
     * @return Integer
     */
    function numIslands($grid) {
        if(empty($grid)){
            return 0;
        }
        $lenR = count($grid);
        $lenC = count($grid[0]);
        $index = 0;
        for($r = 0;$r < $lenR;$r++) {
            for($c = 0;$c < $lenC;$c++) {
                if($grid[$r][$c] == '1'){
                    $this->dfs($grid,$r,$c);
                    $index++;
                }
            }
        }
        return $index;
    }

    function dfs(&$grid,$r,$c){
        if(!$this->inArea($grid,$r,$c)){
            return 0;
        }
        if($grid[$r][$c] != '1') {
            return 0;
        }
        // 将遍历过的岛屿标记为'1'
        $grid[$r][$c] = '0';
        $this->dfs($grid,$r,$c-1);
        $this->dfs($grid,$r,$c+1);
        $this->dfs($grid,$r-1,$c);
        $this->dfs($grid,$r+1,$c);
        return 1;
    }

    function inArea($grid,$r,$c){

        return 0 <= $r && $r < count($grid) && 0 <= $c && $c < count($grid[0]);
    }
}

$i1 = [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]];
$i2 = [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]];

$obj = new Solution();
$obj->numIslands($i2);
