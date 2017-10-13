package main

import (
	"encoding/json"
	"fmt"
	"bytes"
	"io/ioutil"
	"os"
)


type InsideCommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	errMsg  string `json:"errmsg2"` // 小写字母的不会导出到json
}

type CommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	errMsg  string `json:"errmsg2"` // 小写字母的不会导出到json
	Inside InsideCommonError
	ArrayData []InsideCommonError
	MapData map[int]InsideCommonError
}

func main() {

	fullpath := "1.json"
	testWrite := true
	testRead := true

	if testWrite {
		arraydata := []InsideCommonError{InsideCommonError{ErrCode:10}}
		mapdata := make(map[int]InsideCommonError)
		mapdata[10] = InsideCommonError{ErrCode:10}
		ce := CommonError{10, "errcode=10", "code=10", InsideCommonError{ErrMsg:"inside1"},
						arraydata, mapdata}
		data, err := json.Marshal(&ce)

		if err == nil {
			fmt.Println(string(data))
		}

		if err == nil {
			var out bytes.Buffer
			err = json.Indent(&out, data, "", "\t")
			if err == nil {
				ioutil.WriteFile(fullpath, out.Bytes(), os.ModePerm)
			}
		}
	}
	if testRead {
		ce := CommonError{}
		data, err := ioutil.ReadFile(fullpath)
		if err != nil {
			return
		}
		json.Unmarshal(data, &ce)
		fmt.Println(ce)
	}


}
