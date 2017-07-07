package main

import (
	"fmt"
	"unsafe"
)

/*
	学习：变量，const，枚举，类型转换

	需要用var声明变量
	变量名在前，类型在后，比如 var a int
	:= 也可以声明并赋值变量，但是只能在函数内使用
	函数内，声明的变量如果没有使用过，会报错。
	_ 表示占位符（解决变量声明，但没有使用的问题）
	golang没有枚举类型，可以用const代替
*/

var a int = 1
var b int
var c, d, e float32
var aa, bb, cc = 1, "bb", 0.5

// 错误的
// dd := 10

// 常量
const ca, cb, cd = 1, "cb", 0.5
const (
	ce, cf      = "ce", 10
	cg     bool = false
)
const (
	cca = "abc"
	//ccb = len(a)
	ccc = unsafe.Sizeof(cca)
)
const (
	ccd byte = 100
	// 下面这个错误了，因为溢出了。1e20表示1后面20个0
	// cce int  = 1e20
)

// golang没有枚举，但是可以这么定义枚举
const (
	Sunday = iota // 0
	Monday        // 1
	Tuesday
	Wesnesday
	Thursday
	Friday
	Saturday
)

// 没有iota之后，数值需要自己制定，否则常量值后一个等于前一个的值
const (
	eSunday    = 0 // 0
	eMonday        // 0
	eTuesday   = 5 // 5
	eWesnesday     //  5
	eThursday      //  5
	eFriday        //  5
	eSaturday      //  5
)

// 枚举还能这样
const (
	eA, eB = iota, iota << 10 // 0, 0 << 10
	eC, eD                    // 1, 1 << 10
)

// 枚举还能这样
const (
	_        = iota             // iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1
	MB                          // 与KB 表达式相同，但iota = 2
	GB
	TB
)

func main() {

	test_conv()

	ee, ff, gg := 1, "ff", 0.5
	fmt.Println("hello", ee, ff, gg)
	// 变量内存地址
	fmt.Println("address:", &ee)

	// 函数体内变量没用过，报错
	// var aaa int
	// ccc := 10

	array, i := [3]int{0, 1, 2}, 0
	// 多变量赋值时，从左到右依次进行赋值
	i, array[i] = 2, 100
	fmt.Println(array) // 结果是 [100 1 2]

	x, y := "hello", 20
	fmt.Println(x, y)
	{
		// 代码段内重新定义的变量与代码段外的变量没有任何关系
		x, y := "world", 30
		fmt.Println(x, y)
	}
	fmt.Println(x, y) // 输出 hello 20
	{
		// 代码段内赋值的变量会起作用
		x, y = "world2", 40
	}
	fmt.Println(x, y) // 输出 world2 40

	// 没有使用的const变量，不会报错
	const cca = "unused"
	// 此处的const ca与外面的没影响。
	const ca = "ca"

	fmt.Println("Saturday:", Saturday, eSaturday)
}

/*
	基础数据类型
	bool, byte(uint8)
	rune (长度4字节)
	int,uint （长度4字节或者8字节）
	int8,uint8,int16,uint16,int32,uint32,int64,uint64
	float32, float64
	complex64（8字节）
	complex128（16字节）
	uintptr(长度4字节或者8字节)
*/
