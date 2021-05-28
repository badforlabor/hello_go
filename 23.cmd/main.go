/**
 * Auth :   liubo
 * Date :   2020/7/6 17:23
 * Comment: 命令行
 */

package main

import (
	"fmt"
	"os/exec"
)

func main() {

	if err := exec.Command("cmd", "/C", "shutdown", "/r", "/f", "/t", "120").Run(); err != nil {
		fmt.Println("重启电脑失败: ", err.Error())
	}

}
