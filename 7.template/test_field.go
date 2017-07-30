package main

import (
	"os"
	"text/template"
)

type Person1 struct {
	UserName string
	email    string  // 不是导出的，最终不会应用到template中。
}

func testField() {

	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}! {{.email}}")
	p := Person1{UserName: "Astaxie"}
	t.Execute(os.Stdout, p)

}
