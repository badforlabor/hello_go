package main

import "fmt"

/*
	各种函数
	支持多返回值
	支持匿名函数和闭包
	defer 延迟调用
	巧用defer，配合recover函数可以捕获异常（用panic抛异常，或者空指针，数组越界异常）
*/

// 多个返回值
func fun1(x, y int, s string) (int, string) {
	return x + y, s
}

// 多个返回值的另一种形式
func fun11(x, y int, s string) (sum int, ss string) {
	sum = x + y
	ss = s
	return
	// 或者
	// return sum, ss
}

// 函数作为参数
func fun2(fn func() int) int {
	return fn()
}

// 将函数转义成一种类型，这样看着方便
type TypedFunc func(x, y int, s string) (int, string)

func fun3(tf TypedFunc, x, y int, s string) (int, string) {
	return tf(x, y, s)
}

// 变参，只能有一个，且放在最后
func fun4(s string, n ...int) string {
	var sum int
	for _, v := range n {
		sum += v
	}
	return fmt.Sprintf(s, sum)
}

func deferAdd(x, y int) (z int) {

	// 匿名函数，延迟执行
	defer func() {
		z += 100
	}()

	// 匿名函数可以直接赋值给函数变量
	fn := func() {
		z += 0
	}
	defer fn()

	z = x + y
	return
}

// 匿名函数作为变量（当然，这样看着比较费解，最还用type转义一下）
var fns = [](func(x int) int){
	func(x int) int { return x + 1 },
	func(x int) int { return x + 2 },
}

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

func main() {

	// 捕获所有异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	trycatch()

	// 显示103
	fmt.Println(deferAdd(1, 2))

	func() { panic("123") }()

	fmt.Println("test end")
}
