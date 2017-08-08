package main

import (
	"fmt"
	"net/http"
)

func main() {


	err := http.ListenAndServe(":80", nil) //设置监听的端口
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}

	fmt.Println("test end")
}
