package main

import (
	"encoding/json"
	"fmt"
)

type CommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	errMsg  string `json:"errmsg2"` // 小写字母的不会导出到json
}

func main() {
	ce := CommonError{10, "errcode=10", "code=10"}
	data, err := json.Marshal(&ce)

	if err == nil {
		fmt.Println(string(data))
	}

}
