<?php

// 827. Making A Large Island

// You are given an n x n binary matrix grid. You are allowed to change at most one 0 to be 1.
//
//Return the size of the largest island in grid after applying this operation.
//
//An island is a 4-directionally connected group of 1s.

// 思路来自leetcode-cn

// 1.先把岛屿分别使用编号标记，编号起始为2
// 2.将标记好编号的岛屿对应面积记录到map中
// 3.遍历海洋，查找其四周的岛屿
// 4.如果一块海洋附近有岛屿，则记录其编号
// 5.将这块海洋附近的岛屿面积想加+1，和岛屿最大面积比较
class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function largestIsland($grid) {
        if(empty($grid)){
            return 0;
        }
        $index = 1; // 岛屿索引编号初始值
        $maxArea = 0;
        $islandArea = [];
        $lenR = count($grid);
        $lenC = count($grid[0]);
        for($r = 0;$r < $lenR;$r++){
            for($c = 0;$c < $lenC;$c++){
                if($grid[$r][$c] == 1){
                    $index++;
                    $area = $this->island($grid,$r,$c,$index);
                    // 保存当前岛屿编号 -> 面积
                    $islandArea[$index] = $area;
                    // 记录最大岛屿面积
                    $maxArea = max($area,$maxArea);
                }
            }
        }
        // 如果没有陆地，则造一块
        if($maxArea == 0){
            return 1;
        }
        // 寻找海洋附近岛屿
        for($r = 0;$r < $lenR;$r++){
            for($c = 0;$c < $lenC;$c++){
                if($grid[$r][$c] == 0){

                    $oceanIsland = $this->ocean($grid,$r,$c);
                    if(empty($oceanIsland)){
                        continue;
                    }
                    // 遍历海洋四周岛屿信息，获取其面积
                    $island = 1;
                    foreach ($oceanIsland as $ko => $vo){
                        $island += $islandArea[$ko];
                    }
                    // 对比岛屿最大面积和海洋四周岛屿最大面积+1
                    $maxArea = max($island,$maxArea);
                }
            }
        }
        return $maxArea;
    }

    /**
     * 查找海洋格子四周，如果是岛屿，则记录其编号
     * @param $grid
     * @param $r
     * @param $c
     * @return array|mixed
     */
    function ocean($grid,$r,$c) {
        $oceanIsland = []; // 当前海洋格子四周岛屿[岛屿编号：1]
        if($this->inArea($grid,$r,$c-1) && $grid[$r][$c-1] != 0){
            $oceanIsland[$grid[$r][$c-1]] = 1;
        }
        if($this->inArea($grid,$r,$c+1) && $grid[$r][$c+1] != 0){
            $oceanIsland[$grid[$r][$c+1]] = 1;
        }
        if($this->inArea($grid,$r-1,$c) && $grid[$r-1][$c] != 0){
            $oceanIsland[$grid[$r-1][$c]] = 1;
        }
        if($this->inArea($grid,$r+1,$c) && $grid[$r+1][$c] != 0){
            $oceanIsland[$grid[$r+1][$c]] = 1;
        }
        return $oceanIsland;
    }

    function island(&$grid,$r,$c,$index){
        if (!$this->inArea($grid,$r,$c)) {
            return 0;
        }
        if($grid[$r][$c] != 1) {
            return 0;
        }
        // 标记岛屿编号，并返回其面积
        $grid[$r][$c] = $index;
        return 1 + $this->island($grid,$r,$c-1,$index) +
            $this->island($grid,$r,$c+1,$index) +
            $this->island($grid,$r-1,$c,$index) +
            $this->island($grid,$r+1,$c,$index);
    }

    function inArea($grid,$r,$c){

        return 0 <= $r && $r < count($grid) && 0 <= $c && $c < count($grid[0]);
    }
}

$island32 = [
    [0,0,1,0,0,0,0,1,0,0,0,0,0],
    [0,0,0,0,0,0,0,1,1,1,0,0,0],
    [0,1,1,0,1,0,0,0,0,0,0,0,0],
    [0,1,0,0,1,1,0,0,1,0,1,0,0],
    [0,1,0,0,1,1,0,0,1,1,1,0,0],
    [0,0,0,0,0,0,0,0,0,0,1,0,0],
    [0,0,0,0,0,0,0,1,1,1,0,0,0],
    [0,0,0,0,0,0,0,1,1,0,0,0,0],
];

$island33 = [
    [1, 0],
    [0, 1],
];

$island34 =  [
    [1, 1], [1, 0],
];

$island35 = [
    [1, 1],
    [1, 1],
];

$obj = new Solution();
$ret = $obj->largestIsland($island35);
var_dump($ret);