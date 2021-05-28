/**
 * Auth :   liubo
 * Date :   2020/7/1 17:34
 * Comment: 强制类型转换
 */

package main

import "fmt"

type CastA struct {

}
type CastB struct {
	b bool
}

func testCast() {

	{
		var a interface{}
		a = nil
		var b, ok = a.(int)
		if ok {
			fmt.Println(b)
		} else {
			fmt.Println("转化失败:", ok)
		}

	}

	{
		var a interface{} = &CastA{}
		var b = a.(*CastB)	// 这种会直接崩溃的
		if b != nil {
			fmt.Println(b.b)
		}
	}
}