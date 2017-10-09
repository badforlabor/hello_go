package main

import (
	"fmt"
	"reflect"
)

/*
	反射
		参考：http://www.cnblogs.com/coder2012/p/4881854.html
*/

type T struct {
	A int
	B string
	c float32
	tagValue int `tag1:"1" tag2:"2"`
}

func main() {

	testAutoConfig()
	return

	t := T{1,"2",3, 4}

	s := reflect.ValueOf(&t).Elem()
	typeofT := s.Type()
	for i:=0; i<s.NumField(); i++ {
		f := s.Field(i)
		var value interface{}
		func() {
			defer func() {
				if err := recover(); err != nil {
					value = "nil"
				}
			}()
			value = f.Interface()
		}()
		fmt.Printf("idx=%d, variable-name:%s, variable-type:%s variable-value:%s",
				i, typeofT.Field(i).Name, typeofT.Field(i).Type.String(), value)
		fmt.Printf(", tag:%s, tag[tag1]", typeofT.Field(i).Tag, typeofT.Field(i).Tag.Get("tag1"))
		fmt.Printf("\n")
	}


	fmt.Println("test end")
}
