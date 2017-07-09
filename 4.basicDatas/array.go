package main

import (
	"fmt"
	"unsafe"
)

/*
	数组类型是值类型，赋值和传参会复制整个数组！（跟C语言不同，C语言传参传的是指针）
	数组长度必须是常量（即固定的），[2]int和[3]int是不同类型。
	支持 ==，!= 操作符
	指针数组 [n]*T，数组指针 *[n]T
	用len获取数组长度
	获取数组地址的方法：&array，也可以用&array[0]
	[]int{1,2} 和 [...]int{1,2} 不一样！前者是slice，后者是array
*/

func testArray() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	d := []int{1, 2, 3, 4}

	// 这个某个索引的值，比如c[2]=100, c[4]=200
	c := [5]int{2: 100, 4: 200}

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)

	// 结构体数组，注意逗号，得多一个
	// 下面这个数组长度为2
	e := [...]struct {
		name string
		age  int8
	}{
		{"user1", 10},
		{"user2", 20}, // 这里多个逗号，有些奇葩
	}
	fmt.Println("e:", e, len(e), unsafe.Sizeof(e), unsafe.Sizeof(e[0].name))

	params := [2]int{0, 1}
	testArrayAddr(params)
	testPtrAddr(params)
	fmt.Printf("%p %p\n", &params, &params[0])
}

func testArrayAddr(x [2]int) {
	fmt.Printf("%p %p\n", &x, &x[0])
}

// 如果传递指针？
func testPtrAddr(x interface{}) {
	//fmt.Printf("%p\n", x)
}
