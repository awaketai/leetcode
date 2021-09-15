package main

import (
	"fmt"
	"strings"
)

// 917. Reverse Only Letters

// Given a string s, reverse the string according to the following rules:
//
//All the characters that are not English letters remain in the same position.
//All the English letters (lowercase or uppercase) should be reversed.
//Return s after reversing it.

// ascii A-Z 65-90
// a-z 97-122
func main()  {
	//s1 := "Test1ng-Leet=code-Q!"
	s2 := "a-bC-dEf-ghIj"
	ret := reverseOnlyLetters(s2)
	fmt.Printf("%s",ret)
}
// 一个接一个输出s的所有字符，当遇到一个字母时，逆序找到下一个字母
func reverseOnlyLetters(s string) string {
	length := len(s)
	end := length - 1
	arr := make([]string,0)
	for i := 0;i < length;i++ {
		if isLetter(s[i]) {
			// 逆序找到下一个字母
			for !isLetter(s[end]) {
				end--
			}
			arr = append(arr,string(s[end]))
			end--
		}else{
			arr = append(arr,string(s[i]))
		}
	}

	return strings.Join(arr,"")
}

func isLetter(i uint8) bool {
	if (65 <= i && i <= 90) || (97 <= i && i <= 122) {
		return true
	}
	return false
}
