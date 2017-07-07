package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

type wxShareParams struct {
	Appid, Nonce, Signature string
	Timestamp               int64
}

const M = `姓名：{{.Appid}}  
级别：{{.Nonce}}  
性别：{{.Signature}}
分数: [{{.Timestamp}}]
`

func main() {
	timestamp := time.Now().Unix()
	info := wxShareParams{"曦晨", "1", "男", timestamp}
	tm := template.New("")
	tm.Parse(M)
	tm.Execute(os.Stdout, info)

	t, _ := template.ParseFiles("share.gtpl")
	t.Execute(os.Stdout, info)

	fmt.Println(info)
}
