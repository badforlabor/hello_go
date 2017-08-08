package main

import (
	"fmt"
)

/*
	参考：
	http://studygolang.com/articles/8023
	http://blog.csdn.net/sryan/article/details/52353937
	https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html
	http://www.jianshu.com/p/bee02c18b221
*/
/*
	重要：
		{{.}} 表示取当前变量
*/


func main() {

	testField()
	testNested()
	testFunc()
	testIf()


	fmt.Println("test end")
}
