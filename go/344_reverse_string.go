package main

import "fmt"

// 344. Reverse String

// Write a function that reverses a string. The input string is given as an array of characters s.

// Example 1:
//
//Input: s = ["h","e","l","l","o"]
//Output: ["o","l","l","e","h"]
//Example 2:
//
//Input: s = ["H","a","n","n","a","h"]
//Output: ["h","a","n","n","a","H"]

func main()  {
	//s1 := []byte{'h','e','l','l','o'}
	s2 := []byte{'H','a','n','n','a','h'}
	reverseString(s2)
}
func reverseString(s []byte)  {
	start,end := 0,len(s) - 1
	for start < end {
		s[start],s[end] = s[end],s[start]
		start++
		end--
	}
	fmt.Printf("%c",s)
}

