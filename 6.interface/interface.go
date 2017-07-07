package main

import "fmt"

/*
	接⼝命名习惯以er 结尾，
	接⼝只有⽅法签名，没有
	接⼝没有数据字段。
	可在接⼝中嵌⼊其他接⼝
	类型可实现多个接⼝。
	interface{} 相当于c#中的object，可以代表类型
	interface{} 转化成其他类型的方法是：比如转化成int，那么这么写 v.(int)
*/

type Stringer interface {
	String() string
}

type Printer interface {
	// 可以嵌套另一个interface
	Stringer
	Print()
}

func Print(v interface{}) {
	fmt.Printf("type:%T, value:%v\n", v, v)

	if i, ok := v.(int); ok {
		fmt.Println("转化成int成功：", i)
	}
}

func Print2(vargs ...interface{}) {
	// 参数列表
}
func nothing(v interface{}) {

}

// 接口可以做匿名变量
type Tester struct {
	s interface {
		String() string
	}
}

func main() {
	Print(1)
	Print("hello world")

	var s Stringer
	var p Printer
	s = p
	// 子集接口无法赋值给超集接口
	//p = s
	nothing(s)
}
