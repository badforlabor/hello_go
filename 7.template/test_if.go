package main

import (
	"os"
	"text/template"
)

/*

*/

func testIf() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} will not be outputted. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("Not empty pipeline if demo: {{if `anything`}} will be outputted. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if part {{else}} else part.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)


	tIfElse = template.New("template test")
	tIfElse = tIfElse.Funcs(template.FuncMap{"funcIf": funcIf})
	tIfElse = template.Must(tIfElse.Parse("{{$b := 10}} {{$a := funcIf $b}}if-else demo: {{if $a}} if part {{else}} else part.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}

func funcIf(args ...interface{}) bool {
	return true
}
