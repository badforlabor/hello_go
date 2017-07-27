package main

import (
	"fmt"
	"reflect"
)

type structA struct {
	v int
}

func main() {
	fmt.Println("123")

	reflectMap := map[string]reflect.Type{
		"structA": reflect.TypeOf(structA{}),
	}
	fmt.Println(reflectMap)

	a := reflect.New(reflectMap["structA"]).Interface().(*structA)

	fmt.Println(a.v, reflect.TypeOf(a), reflect.TypeOf(*a))
}
