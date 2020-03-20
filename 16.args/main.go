/**
 * Auth :   liubo
 * Date :   2019/10/17 9:27
 * Comment: 命令行参数
 */

package main

import (
	"flag"
	"fmt"
	"os"
)

var s = flag.String("msg", "", "show message.")

func main() {
	flag.Parse()

	fmt.Println(os.Args)
	fmt.Println("msg:", *s)
}
