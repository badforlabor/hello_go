package main

import "fmt"

// 巧用defer，配合recover函数可以捕获异常（用panic抛异常，或者空指针，数组越界异常）

func trycatch() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	x := []int{}
	// 这里会抛出一个异常
	fmt.Println(x[0])

	// 手动抛异常
	panic("throw exception.")
}
