/**
 * Auth :   liubo
 * Date :   2020/7/6 10:43
 * Comment: 正则表达式
 */

package main

import (
	"fmt"
	"regexp"
)

func main() {
	var r, _ = regexp.Compile("robot-.*")
	var b = r.MatchString("robot-1123")
	fmt.Println(b)

	var r2, _ = regexp.Compile("")
	fmt.Println(r2 == nil)
}
