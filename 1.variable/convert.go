package main

import "fmt"

/*
	类型转换
	需要显示定义，不支持隐式
	不能将其他类型的变量，当bool使用，比如 if a {}
	使用"``"定义的内容，将不会被转义
*/

func test_conv() {

	{
		var b byte = 100
		// 隐式转化，是错误的
		// var n int = b
		var n int = int(b)
		fmt.Println(n)
	}
	{
		s := `a
		b\r\n\x00
		c`
		fmt.Println(s)
	}
}
