/**
 * Auth :   liubo
 * Date :   2020/3/20 16:35
 * Comment: 事件派发
 */

package main

import (
	"fmt"
	"reflect"
)

type TestA struct {
	A string
}
func (self *TestA) showme(b int) {
	fmt.Println(self.A, b)
}

func main() {

	RegEvent("1", func() {
		fmt.Println("1")
	})
	FireEvent("1")

	var a = &TestA{A:"aaaaa"}
	var b = &TestA{A:"bbbbb"}
	RegEvent("a", a.showme)
	RegEvent("b",  b.showme)

	FireEvent("a", 1)
	FireEvent("b", 2)
}

func init() {
	eventMap = make(map[string]reflect.Value)
}

var eventMap map[string]reflect.Value
func RegEvent(key string, callback interface{}) {
	eventMap[key] = reflect.ValueOf(callback)
}
func FireEvent(key string, args ...interface{}) {
	v, ok := eventMap[key]
	if ok {
		var argsList []reflect.Value
		for _, v := range args {
			argsList = append(argsList, reflect.ValueOf(v))
		}
		v.Call(argsList)
	}
}

