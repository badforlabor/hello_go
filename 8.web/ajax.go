package main

import (
	"net/http"
	"fmt"
	"strings"
	"encoding/json"
)


type CommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func init() {
	http.HandleFunc("/ajax", ajax)
}

func ajax(w http.ResponseWriter, r *http.Request) {

	// 需要设置返回内容是json，要不会当成txt解析
	w.Header().Set("Content-type", "application/json")
	// 需要设置这个，要不会提示错误：No 'Access-Control-Allow-Origin' header is present on the requested resource
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("wx path", strings.Join([]string{"http://", r.Host, r.RequestURI}, ""))

	ret := &CommonError{1, "ok"}
	data, err := json.Marshal(ret)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(data)
	fmt.Println("ajax ret ok")
}



