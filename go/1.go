package main

import "fmt"

//type Student struct {
//	Name string
//	Age int
//}
//
//func StuReg(name string,age int) *Student {
//	s := new(Student) // 局部变量逃逸到堆
//	s.Name = name
//	s.Age = age
//	return s
//}

//func Sli()  {
//	s := make([]int,10000,10000)
//	for i,_ := range  s {
//		s[i] = i
//	}
//}

func Fi() func() int {
	a,b := 0,1
	return func() int {
		a,b = b,a+b
		return a
	}
}


func main()  {
	f := Fi()

	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci: %d\n", f())
	}
}
