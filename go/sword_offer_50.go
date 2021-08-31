package main

import "fmt"

//剑指 Offer 50. 第一个只出现一次的字符

// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
//
//示例:
// s = "abaccdeff"
//返回 "b"
//
//s = ""
//返回 " "

func main()  {
	//s1 := "leetcode"
	//s2 := "abaccdeff"
	s3 := "z"
	ret := firstUniqChar(s3)
	fmt.Println(ret)
}

// 1.按位，数组的每个索引都是字符ascii码值，值是字符出现次数
// 2.遍历字符s，查看每个字符的ascii码在数组中出现的次数，第一个为1的返回
func firstUniqChar(s string) byte {
	var mem = make([]int,123)
	for _,v := range s {
		mem[v]++
	}
	for i,v := range s {
		if mem[v] == 1 {
			return s[i]
		}
	}
	return ' '
}
