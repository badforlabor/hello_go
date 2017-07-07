package main

import "fmt"

/*
	if语句

*/
func testFor() {

	s := "abc"

	// 表达式也不需要括号
	for i, n := 0, len(s); i < n; i++ {
		fmt.Println(s[i])
	}

	n := len(s)
	// while的代替写法
	for n > 0 {
		n--
		fmt.Println(s[n])
	}
}
