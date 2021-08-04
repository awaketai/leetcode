<?php

// 463. Island Perimeter
// You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.
//
//Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).
//
//The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island. One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.


class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function islandPerimeter($grid) {
        if(empty($grid)){
            return 0;
        }
        $lenR = count($grid);
        $lenC = count($grid[0]);
        for($r = 0;$r < $lenR;$r++){
            for($c = 0;$c < $lenC;$c++){
                if($grid[$r][$c] == 1){

                    return $this->dfs($grid,$r,$c);
                }
            }
        }
        return 0;
    }

    function dfs(&$grid,$r,$c){
        // 如果到达边界，则认为触碰到了岛屿一边[黄色]
        if(!$this->inArea($grid,$r,$c)){
            return 1;
        }
        // 如果到达海洋，则认为触碰到了海洋一边[蓝色]
        if($grid[$r][$c] == 0){
            return 1;
        }
        // 如果是已标记过的陆地格子
        if($grid[$r][$c] != 1){
            return 0;
        }
        $grid[$r][$c] = 2;
        // 将遍历过的陆地格子标记为2，避免重复遍历
        return $this->dfs($grid,$r,$c-1) +
            $this->dfs($grid,$r,$c+1) +
            $this->dfs($grid,$r-1,$c) +
            $this->dfs($grid,$r+1,$c);
    }

    function inArea($grid,$r,$c){

        return 0 <= $r && $r < count($grid) && 0 <= $c && $c < count($grid[0]);
    }
}

$p1 = [
    [0,1,0,0],
    [1,1,1,0],
    [0,1,0,0],
    [1,1,0,0],
];

$p2 = [
    [1],
];

$p3 = [
    [1,0],
];
$p4 = [
    [1,1],
    [1,1],
];

$obj = new Solution();
$ret = $obj->islandPerimeter($p4);
var_dump($ret);

