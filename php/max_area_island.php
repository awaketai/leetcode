<?php

// 695. Max Area of Island

// You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
//
//The area of an island is the number of cells with a value 1 in the island.
//
//Return the maximum area of an island in grid. If there is no island, return 0.

$island1 = [
    [1,1,1,1,0],
    [1,1,0,1,0],
    [1,1,0,0,0],
    [0,0,0,0,0],
];
$island2 = [
    [1,1,0,0,0],
    [1,1,0,0,0],
    [0,0,1,0,0],
    [0,0,0,1,1],
];

$island3 = [
    [0,0,1,0,0,0,0,1,0,0,0,0,0],
    [0,0,0,0,0,0,0,1,1,1,0,0,0],
    [0,1,1,0,1,0,0,0,0,0,0,0,0],
    [0,1,0,0,1,1,0,0,1,0,1,0,0],
    [0,1,0,0,1,1,0,0,1,1,1,0,0],
    [0,0,0,0,0,0,0,0,0,0,1,0,0],
    [0,0,0,0,0,0,0,1,1,1,0,0,0],
    [0,0,0,0,0,0,0,1,1,0,0,0,0],
];

$island4 = [
    [0,0,0,0,0,0,0,0],
];

$island5 = [
    [1,1,0,0,0],
    [1,1,0,0,0],
    [0,0,0,1,1],
    [0,0,0,1,1],
];


class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function maxAreaOfIsland($grid) {
        $res = 0;
        $rLen = count($grid);
        $cLen = count($grid[0]);
        for($r = 0;$r < $rLen;$r++){
            for ($c = 0;$c < $cLen;$c++){
                // 如果当前是岛屿，则查看其周围是否也是岛屿
                if($grid[$r][$c] == 1){
                    $a = $this->area($grid,$r,$c);
                    $res = max($res,$a);
                }
            }
        }
        return $res;
    }

    function area(array &$grid,$r,$c) {
        // 如果不在网格内
        if(!$this->inArea($grid,$r,$c)){
            return 0;
        }
        // 如果不是岛屿(1)
        if($grid[$r][$c] != 1){
            return 0;
        }
        // 如果当前岛屿已经遍历过，则将其置为2，避免重复遍历
        $grid[$r][$c] = 2;
        // 没遍历一个格子，如果是岛屿，则面积+1
        return 1 +
            $this->area($grid,$r-1,$c) +
            $this->area($grid,$r,$c-1) +
            $this->area($grid,$r+1,$c) +
            $this->area($grid,$r,$c+1);
    }
    /**
     * 判断当前坐标是否在网格内
     * @param array $grid
     * @param $r int 横坐标
     * @param $c int 竖坐标
     * @return bool
     */
    function inArea(array $grid,$r,$c) {

        // count(grid):行长度
        // count(grid[0]):列长度
        return 0 <= $r && $r < count($grid) && 0 <= $c && $c < count($grid[0]);
    }
}

$obj = new Solution();
$ret = $obj->maxAreaOfIsland($island1);
//var_dump($island1);
if($ret != 9){
    exit("test error-case:island1:value:".$ret." expected:9");
}
$ret = $obj->maxAreaOfIsland($island2);
if($ret != 4){
    exit("test error-case:island2:value:".$ret." expected:4");
}
$ret = $obj->maxAreaOfIsland($island3);
if($ret != 6){
    exit("test error-case:island3:value:".$ret." expected:6");
}
$ret = $obj->maxAreaOfIsland($island4);
if($ret != 0){
    exit("test error-case:island4:value:".$ret." expected:0");
}
$ret = $obj->maxAreaOfIsland($island5);
if($ret != 4){
    exit("test error-case:island5:value:".$ret." expected:4");
}
echo "test success!";