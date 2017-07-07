package main

import "fmt"

/*
	各种运算符，各种类型的语句
	golang没有while关键字，但是可以用for代替
*/
func main() {
	testOp()
	testIf()
	testFor()
	testRange()
	testSwitch()
	testOther()

	fmt.Println("test end")
}
