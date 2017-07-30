/*
	template function
	语法规则：
		{{funcName .param1 .param2}}
		执行Parse或者ParseFiles之前，需要执行先注册函数：t = t.Funcs(template.FuncMap{"writeVariable": WriteVariable})
*/

package main

import (
"fmt"
"html/template"
"os"
"strings"
)

//type Friend struct {
//	Fname string
//}
//
//type Person struct {
//	UserName string
//	Emails   []string
//	Friends  []*Friend
//}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace the @ by " at "
	return (substrs[0] + " at " + substrs[1])
}

func testFunc() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
                {{range .Emails}}
                    an emails {{.|emailDeal}}
                    an emails: {{emailDeal .}}
                {{end}}
                {{with .Friends}}
                {{range .}}
                    my friend name is {{.Fname}}
                {{end}}
                {{end}}
                `)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
