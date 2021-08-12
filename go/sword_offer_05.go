package main

import "fmt"

// 剑指 Offer 05. 替换空格
//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

var s = "We are happy."
var rStr = "%20"

func main()  {
	ret := replaceSpace2(s)
	fmt.Println(ret)
}

func replaceSpace(s string) string {
	var str = ""
	for _,v := range s {
		if string(v) == " " {
			str = str + rStr
		} else {
			str = str + string(v)
		}
	}
	return str
}

func replaceSpace2(s string) string  {
	bStr := []byte(s)
	res := make([]byte,0)
	length := len(bStr)
	for i := 0;i < length;i++ {
		if bStr[i] == ' ' {
			res = append(res,[]byte("%20")...)
		}else{
			res = append(res,bStr[i])
		}
	}
	return string(res)
}
