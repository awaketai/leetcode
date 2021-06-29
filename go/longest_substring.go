package main

import "fmt"

var a,b,c,d string = "abcabcbb","bbbbb","pwwkew",""

func main()  {
	r := lengthOfLongestSubstring(d)
	fmt.Println(r)
}

func lengthOfLongestSubstring(str string) int {
	m := map[byte]int{} // 记录每个字符是否出现过
	n := len(str)
	rk,ans := -1,0
	for i := 0;i < n;i++ {
		if i != 0 {
			delete(m,str[i-1])
		}
		for rk + 1 <n && m[str[rk + 1]] == 0{
			// 不断的移动右指针
			m[str[rk + 1]]++
			rk++
		}
		//第i到rk个字符是一个极长的无重复字符子串
		ans = max(ans,rk - i + 1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
