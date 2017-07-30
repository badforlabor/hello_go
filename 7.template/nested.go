package main

import (
	"os"
	"text/template"
)

/*
	嵌套:
		{{.}} 表示取当前变量
		所以当执行
            {{range $i, $_ := .Emails}}
                email[{{$i}}] = {{.}}		// range取出来第二个变量是匿名的，所以可以用{{.}}代替
            {{end}}


        {{with .UserName}} 相当于把当前变量设置成了 .UserName，这样就可以直接用{{.}}访问了。
			{{with .UserName}}
			my name is {{.}}
			{{end}}

*/

type Friend struct {
	Fname string
	Tags []int
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}


func testNested() {

	f1 := Friend{Fname: "minux.ma", Tags: []int{0,1,2}}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`
	test nested. hello {{.UserName}}!
            {{range .Emails}}
                an email {{.}}
            {{end}}
            {{with .Friends}}
            {{range .}}
                my friend name is {{.Fname}} [{{range .Tags}}{{.}}{{end}}]
            {{end}}
            {{end}}

            {{range $i, $v := .Emails}}
                email[{{$i}}] = {{$v}}
            {{end}}

            {{range $i, $_ := .Emails}}
                email[{{$i}}] = {{.}}
            {{end}}

		{{with .UserName}}
		my name is {{.}}
		{{end}}
            `)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)

}
