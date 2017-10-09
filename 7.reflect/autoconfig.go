package main

import (
	"reflect"
	"fmt"
	"encoding/json"
	"strconv"
	"bytes"
	"io/ioutil"
	"os"
)

/*
	自动config
	不用json序列化，麻烦！
	使用方法很简单，参考testAutoConfig
*/

// 暂时支持基础类型 int,float,string
type autoConfig struct {
	A int
	B string
	C float32
	D int64
	E int8
	F uint8

	nc float32
	ND []int
}

func isInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	}
	return false
}
func isUInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	}
	return false
}
func IsInt(v reflect.Value) bool {
	return isInt(v) || isUInt(v)
}
func IsFloat(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	}
	return false
}
func IsString(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return true
	}
	return false
}
func SetValue(v reflect.Value, vstr string) {

	if !v.CanSet() {
		return
	}

	if isInt(v) {
		i, _ := strconv.ParseInt(vstr, 10, 64)
		v.SetInt(i)
	} else if IsFloat(v) {
		i, _ := strconv.ParseFloat(vstr, 64)
		v.SetFloat(i)
	} else if IsString(v) {
		v.SetString(vstr)
	} else if isUInt(v) {
		i, _ := strconv.ParseUint(vstr, 10, 64)
		v.SetUint(i)
	}
}

// 遍历反射内容
func loopValue(objPtr interface{}, callback func(int, reflect.StructField, reflect.Value)) {
	s := reflect.ValueOf(objPtr).Elem()
	typeofT := s.Type()
	for i:=0; i<s.NumField(); i++ {
		variableName := typeofT.Field(i).Name
		if variableName[0] >= 'A' && variableName[0] <= 'Z' {
		} else {
			// 大写字母开头的才序列化
			continue
		}
		// 只支持基础数据类型
		if IsInt(s.Field(i)) || IsFloat(s.Field(i)) || IsString(s.Field(i)) {

		} else {
			continue
		}
		callback(i, typeofT.Field(i), s.Field(i))
	}

}

func SaveAutoConfig(objPtr interface{}, fullpath string) {
	jsonmap := make(map[string]string)

	loopValue(objPtr, func(idx int, field reflect.StructField, value reflect.Value) {
		jsonmap[field.Name] = fmt.Sprint(value.Interface())
	})

	datas, err := json.Marshal(jsonmap)
	if err == nil {
		var out bytes.Buffer
		err = json.Indent(&out, datas, "", "\t")
		if err == nil {
			ioutil.WriteFile(fullpath, out.Bytes(), os.ModePerm)
		}
	}
}
func LoadAutoConfig(objPtr interface{}, fullpath string) {

	data, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return
	}

	jsonmap := make(map[string]string)
	json.Unmarshal(data, &jsonmap)

	loopValue(objPtr, func(idx int, field reflect.StructField, value reflect.Value) {
		if vstr, ok := jsonmap[field.Name]; ok {
			SetValue(value, vstr)
		}
	})
}

func testAutoConfig() {
	t := autoConfig{1,"c:\\1/2.txt", 3, 4, 5, 6, 7, []int{1,2}}
	SaveAutoConfig(&t, "1.json")

	t2 := autoConfig{}
	fmt.Println("before:", t2)
	LoadAutoConfig(&t2, "1.json")
	fmt.Println("after:", t2)
}
