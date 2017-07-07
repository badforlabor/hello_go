package main

import "fmt"

/*
	range语句：返回 （索引，值）或者（键，值）
		支持：string, array, slice, map
	相当于iterator
*/
func testRange() {

	// string
	s := "abc"

	// 正统写法
	for i, v := range s {
		fmt.Println("i:", i, "v:", v)
	}

	// 忽略第二个变量
	for i := range s {
		fmt.Println(s[i])
	}

	// 忽略第一个变量
	for _, v := range s {
		fmt.Println("v:", v)
	}

	// 两个返回值都忽略
	for range s {
		fmt.Println("v:")
	}

	// range map用法
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
