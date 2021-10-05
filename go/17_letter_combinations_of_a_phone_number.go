package main

import (
	"context"
	"github.com/gpmgo/gopm/modules/cli"
)

// 17. Letter Combinations of a Phone Number

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
//
//A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

// 回溯解法

var digits = map[string]string{
	"2":"abc",
	"3":"def",
	"4":"ghi",
	"5":"jkl",
	"6":"mno",
	"7":"pqrs",
	"8":"tuv",
	"9":"wzyx",
}
var combinations []string
func letterCombinations(digits string) []string {

	length := len(digits)
	if length == 0 {
		return []string{}
	}

	combinations = make([]string,0)

}

func backtrack(digits,combination string,index int)  {
	if index == len(digits) {
		combinations = append(combinations,combination)
	}else{
		digit := string(digits[index])
	}
}